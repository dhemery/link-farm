package main

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"
)

type pathRuleTest struct {
	FS   fs.FS
	Path string
	Want error
}

var packagePathRuleTests = map[string]pathRuleTest{
	"path to readable dir is good": {
		FS: fstest.MapFS{
			"path/to/readable/dir": directory(0444),
		},
		Path: "path/to/readable/dir",
		Want: nil,
	},
	"path to unreadable dir is permission error": {
		FS: fstest.MapFS{
			"path/to/unreadable/dir": directory(0333),
		},
		Path: "path/to/unreadable/dir",
		Want: fs.ErrPermission,
	},
	"path to link is invalid": {
		FS: fstest.MapFS{
			"path/to/link": linkTo("some/place"),
		},
		Path: "path/to/link",
		Want: fs.ErrInvalid,
	},
	"path to file is invalid": {
		FS: fstest.MapFS{
			"path/to/file": regularFile(),
		},
		Path: "path/to/file",
		Want: fs.ErrInvalid,
	},
	"path to non-existent file is not exist error": {
		FS:   fstest.MapFS{},
		Path: "path/to/non-existent/file",
		Want: fs.ErrNotExist,
	},
}

var installPathRuleTests = map[string]pathRuleTest{
	"path to readable dir is good": {
		FS: fstest.MapFS{
			"path/to/readable/dir": directory(0444),
		},
		Path: "path/to/readable/dir",
		Want: nil,
	},
	"path to unreadable dir is permission error": {
		FS: fstest.MapFS{
			"path/to/unreadable/dir": directory(0333),
		},
		Path: "path/to/unreadable/dir",
		Want: fs.ErrPermission,
	},
	"path to link is invalid": {
		FS: fstest.MapFS{
			"path/to/link": linkTo("some/place"),
		},
		Path: "path/to/link",
		Want: fs.ErrInvalid,
	},
	"path to file is invalid": {
		FS: fstest.MapFS{
			"path/to/file": regularFile(),
		},
		Path: "path/to/file",
		Want: fs.ErrInvalid,
	},
	"path to farm dir is permission error": {
		FS: fstest.MapFS{
			"root/farm-dir/.farm": regularFile(),
		},
		Path: "root/farm-dir",
		Want: fs.ErrPermission,
	},
	"path to dir inside farm is permission error": {
		FS: fstest.MapFS{
			"root/farm-dir/.farm":                   regularFile(),
			"root/farm-dir/dir/dir/dir-inside-farm": directory(0755),
		},
		Path: "root/farm-dir/dir/dir/dir-inside-farm",
		Want: fs.ErrPermission,
	},
}

func TestPackagePathRules(t *testing.T) {
	for name, test := range packagePathRuleTests {
		t.Run(name, func(t *testing.T) {
			got := CheckPackagePath(test.FS, test.Path)
			if !errors.Is(got, test.Want) {
				t.Errorf("got error %v, want %v", got, test.Want)
			}
		})
	}
}

func TestInstallPathRules(t *testing.T) {
	for name, test := range installPathRuleTests {
		t.Run(name, func(t *testing.T) {
			got := CheckInstallPath(test.FS, test.Path)
			if !errors.Is(got, test.Want) {
				t.Errorf("got error %v, want %v", got, test.Want)
			}
		})
	}
}
