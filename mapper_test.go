package main

import (
	"errors"
	"io/fs"
	"path"
	"testing"
	"testing/fstest"
)

type entryFunc func(fsys fstest.MapFS, path string)

type mapperTest struct {
	RelPath     string
	SourceEntry entryFunc
	TargetEntry entryFunc
	Error       error
	Action      Action
}

var mapperTests = map[string]mapperTest{
	"create link if target has no entry at source file path": {
		RelPath:     "path/to/file",
		SourceEntry: regularFile,
		TargetEntry: noEntry,
		Action: CreateLinkAction{
			From: "target-dir/path/to/file",
			To:   "source-dir/path/to/file",
		},
		Error: nil,
	},
	"create link if target has no entry at source dir path": {
		RelPath:     "path/to/dir",
		SourceEntry: emptyDirectory,
		TargetEntry: noEntry,
		Action: CreateLinkAction{
			From: "target-dir/path/to/dir",
			To:   "source-dir/path/to/dir",
		},
		Error: nil,
	},
	"error if target has existing file at source file path": {
		RelPath:     "path/to/file",
		SourceEntry: regularFile,
		TargetEntry: regularFile,
		Action:      nil,
		Error:       fs.ErrExist,
	},
	"replace with link if target has empty dir at source dir path": {
		RelPath:     "path/to/dir",
		SourceEntry: emptyDirectory,
		TargetEntry: emptyDirectory,
		Action: ReplaceWithLink{
			From: "target-dir/path/to/dir",
			To:   "source-dir/path/to/dir",
		},
		Error: nil,
	},
}

func TestMapper(t *testing.T) {
	for name, test := range mapperTests {
		t.Run(name, func(t *testing.T) {
			sourceDir := "source-dir"
			targetDir := "target-dir"
			sourcePath := path.Join(sourceDir, test.RelPath)
			targetPath := path.Join(targetDir, test.RelPath)

			fsys := fstest.MapFS{}
			test.SourceEntry(fsys, sourcePath)
			test.TargetEntry(fsys, targetPath)

			mapper := Mapper{FS: fsys, SourceDir: sourceDir, TargetDir: targetDir}
			action, err := mapper.Map(sourcePath)
			if action != test.Action {
				t.Errorf("got action %#v, want %#v", action, test.Action)
			}
			if !errors.Is(err, test.Error) {
				t.Errorf(`got error %v, want "%v"`, err, test.Error)
			}
		})
	}
}

func emptyDirectory(fsys fstest.MapFS, path string) {
	fsys[path] = &fstest.MapFile{Mode: 0644 | fs.ModeDir}
}

func regularFile(fsys fstest.MapFS, path string) {
	fsys[path] = &fstest.MapFile{Mode: 0644}
}

func noEntry(fsys fstest.MapFS, path string) {
	delete(fsys, path)
}
