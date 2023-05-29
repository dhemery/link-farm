package main

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"
)

type mapperTest struct {
	SourceFS fstest.MapFS
	TargetFS fstest.MapFS
	Path     string
	Action   Action
	Error    error
}

var mapperTests = map[string]mapperTest{
	"link target to source file if target has no entry": {
		Path: "path/to/entry",
		SourceFS: fstest.MapFS{
			"path/to/entry": regularFile(),
		},
		TargetFS: fstest.MapFS{}, // No entry at path
		Action:   CreateLink{Path: "path/to/entry"},
		Error:    nil,
	},
	"link target to source dir if target has no entry": {
		Path: "path/to/entry",
		SourceFS: fstest.MapFS{
			"path/to/entry": directory(),
		},
		TargetFS: fstest.MapFS{}, // No entry at path
		Action:   CreateLink{Path: "path/to/entry"},
		Error:    nil,
	},
	"cannot link existing target file to source file": {
		Path: "path/to/entry",
		SourceFS: fstest.MapFS{
			"path/to/entry": regularFile(),
		},
		TargetFS: fstest.MapFS{
			"path/to/entry": regularFile(),
		},
		Action: nil,
		Error:  fs.ErrExist,
	},
	"cannot link existing target file to source dir": {
		Path: "path/to/entry",
		SourceFS: fstest.MapFS{
			"path/to/entry": directory(),
		},
		TargetFS: fstest.MapFS{
			"path/to/entry": regularFile(),
		},
		Action: nil,
		Error:  fs.ErrExist,
	},
	"cannot link existing target dir to source file": {
		Path: "path/to/entry",
		SourceFS: fstest.MapFS{
			"path/to/entry": regularFile(),
		},
		TargetFS: fstest.MapFS{
			"path/to/entry": directory(),
		},
		Action: nil,
		Error:  fs.ErrExist,
	},
	"descend if source and target are both dirs": {
		Path: "path/to/entry",
		SourceFS: fstest.MapFS{
			"path/to/entry": directory(),
		},
		TargetFS: fstest.MapFS{
			"path/to/entry": directory(),
		},
		Action: Descend{},
		Error:  nil,
	},
}

func TestMapper(t *testing.T) {
	for name, test := range mapperTests {
		t.Run(name, func(t *testing.T) {

			mapper := Mapper{Source: test.SourceFS, Target: test.TargetFS}
			action, err := mapper.Map(test.Path)
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
