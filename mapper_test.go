package main

import (
	"errors"
	"fmt"
	"io/fs"
	"path"
	"testing"
	"testing/fstest"
)

type fileFunc func(fs fstest.MapFS, path string)

type mapperTest struct {
	RelPath     string
	SourceEntry fileFunc
	TargetEntry fileFunc
	Error       error
	Action      Action
}

var mapperTests = map[string]mapperTest{
	"create link if target has no entry at source file path": {
		RelPath:     "path/to/file",
		SourceEntry: regularFile,
		TargetEntry: noEntry,
		Error:       nil,
		Action:      CreateLinkAction{From: "/target-dir/path/to/file", To: "/source-dir/path/to/file"},
	},
	"error if target has existing file at source file path": {
		RelPath:     "path/to/file",
		SourceEntry: regularFile,
		TargetEntry: regularFile,
		Error:       fs.ErrExist,
		Action:      nil,
	},
}

func TestMapper(t *testing.T) {
	for name, test := range mapperTests {
		t.Run(name, func(t *testing.T) {
			sourceDir := "/source-dir"
			targetDir := "/target-dir"
			sourcePath := path.Join(sourceDir, test.RelPath)
			targetPath := path.Join(targetDir, test.RelPath)

			fsys := fstest.MapFS{}
			test.SourceEntry(fsys, sourcePath)
			test.TargetEntry(fsys, targetPath)
			fsys["foo"] = &fstest.MapFile{Mode: 0644}

			for p, f := range fsys {
				stat, err := fsys.Stat(p)
				fmt.Println(name, "PATH", p, "FILE", f, "STAT", stat, "ERR", err)
			}

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

func regularFile(fsys fstest.MapFS, path string) {
	fsys[path] = &fstest.MapFile{Mode: 0644}
	file := fsys[path]
	stat, err := fsys.Stat(path)
	fmt.Println("regularFile", "PATH", path, "FILE", file, "STAT", stat, "ERR", err)
}

func noEntry(fsys fstest.MapFS, path string) {
	delete(fsys, path)
}
