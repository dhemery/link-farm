package main

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"
)

type testLinker struct{}

func (t testLinker) Symlink(oldname, newname string) error {
	return nil
}

type mapperTest struct {
	Mapper     Mapper
	SourcePath string
	Action     Action
	Error      error
}

var mapperTests = map[string]mapperTest{
	"create link to source file if target has no entry": {
		Mapper: Mapper{
			SourceDir: "source-dir",
			TargetDir: "target-dir",
			FS: fstest.MapFS{
				"source-dir/path/to/entry": regularFile(),
				// Target dir has no entry
			},
			Linker: testLinker{},
		},
		SourcePath: "source-dir/path/to/entry",
		Action: CreateLink{
			Linker: testLinker{},
			From:   "target-dir/path/to/entry",
			To:     "source-dir/path/to/entry",
		},
	},
	"create link to source dir if target has no entry": {
		Mapper: Mapper{
			SourceDir: "source-dir",
			TargetDir: "target-dir",
			FS: fstest.MapFS{
				"source-dir/path/to/entry": directory(),
				// Target dir has no entry
			},
			Linker: testLinker{},
		},
		SourcePath: "source-dir/path/to/entry",
		Action: CreateLink{
			Linker: testLinker{},
			From:   "target-dir/path/to/entry",
			To:     "source-dir/path/to/entry",
		},
	},
	"cannot link existing target file to source file": {
		Mapper: Mapper{
			SourceDir: "source-dir",
			TargetDir: "target-dir",
			FS: fstest.MapFS{
				"source-dir/path/to/entry": regularFile(),
				"target-dir/path/to/entry": regularFile(),
			},
		},
		SourcePath: "source-dir/path/to/entry",
		Error:      fs.ErrExist,
	},
	"cannot link existing target file to source dir": {
		Mapper: Mapper{
			SourceDir: "source-dir",
			TargetDir: "target-dir",
			FS: fstest.MapFS{
				"source-dir/path/to/entry": directory(),
				"target-dir/path/to/entry": regularFile(),
			},
		},
		SourcePath: "source-dir/path/to/entry",
		Error:      fs.ErrExist,
	},
	"cannot link existing target dir to source file": {
		Mapper: Mapper{
			SourceDir: "source-dir",
			TargetDir: "target-dir",
			FS: fstest.MapFS{
				"source-dir/path/to/entry": regularFile(),
				"target-dir/path/to/entry": directory(),
			},
		},
		SourcePath: "source-dir/path/to/entry",
		Error:      fs.ErrExist,
	},
	"descend if source and target are both dirs": {
		Mapper: Mapper{
			SourceDir: "source-dir",
			TargetDir: "target-dir",
			FS: fstest.MapFS{
				"source-dir/path/to/entry": directory(),
				"target-dir/path/to/entry": directory(),
			},
		},
		SourcePath: "source-dir/path/to/entry",
		Action:     Descend{},
	},
}

func TestMapper(t *testing.T) {
	for name, test := range mapperTests {
		t.Run(name, func(t *testing.T) {
			action, err := test.Mapper.Map(test.SourcePath)
			if action != test.Action {
				t.Errorf("got action %#v, want %#v", action, test.Action)
			}
			if !errors.Is(err, test.Error) {
				t.Errorf(`got error %v, want "%v"`, err, test.Error)
			}
		})
	}
}

func directory() *fstest.MapFile {
	return &fstest.MapFile{Mode: 0644 | fs.ModeDir}
}

func regularFile() *fstest.MapFile {
	return &fstest.MapFile{Mode: 0644}
}
