package main

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"
)

type mapperTest struct {
	GroveFS fstest.MapFS
	UserFS  fstest.MapFS
	Path    string
	Action  Action
	Error   error
}

var mapperTests = map[string]mapperTest{
	"create link to grove file if target has no entry": {
		Path: "path/to/entry",
		GroveFS: fstest.MapFS{
			"path/to/entry": regularFile(),
		},
		UserFS: fstest.MapFS{}, // No entry at path
		Action: CreateLink{Path: "path/to/entry"},
		Error:  nil,
	},
	"create link to grove dir if target has no entry": {
		Path: "path/to/entry",
		GroveFS: fstest.MapFS{
			"path/to/entry": emptyDirectory(),
		},
		UserFS: fstest.MapFS{}, // No entry at path
		Action: CreateLink{Path: "path/to/entry"},
		Error:  nil,
	},
	"cannot map existing user file to grove file": {
		Path: "path/to/entry",
		GroveFS: fstest.MapFS{
			"path/to/entry": regularFile(),
		},
		UserFS: fstest.MapFS{
			"path/to/entry": regularFile(),
		},
		Action: nil,
		Error:  fs.ErrExist,
	},
	"cannot map existing user file to grove dir": {
		Path: "path/to/entry",
		GroveFS: fstest.MapFS{
			"path/to/entry": emptyDirectory(),
		},
		UserFS: fstest.MapFS{
			"path/to/entry": regularFile(),
		},
		Action: nil,
		Error:  fs.ErrExist,
	},
	"replace empty user dir with link to source dir": {
		Path: "path/to/entry",
		GroveFS: fstest.MapFS{
			"path/to/entry": emptyDirectory(),
		},
		UserFS: fstest.MapFS{
			"path/to/entry": emptyDirectory(),
		},
		Action: ReplaceWithLink{Path: "path/to/entry"},
		Error:  nil,
	},
}

func TestMapper(t *testing.T) {
	for name, test := range mapperTests {
		t.Run(name, func(t *testing.T) {

			mapper := Mapper{Source: test.GroveFS, Target: test.UserFS}
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

func emptyDirectory() *fstest.MapFile {
	return &fstest.MapFile{Mode: 0644 | fs.ModeDir}
}

func regularFile() *fstest.MapFile {
	return &fstest.MapFile{Mode: 0644}
}
