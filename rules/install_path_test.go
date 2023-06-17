package rules

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"
)

type installPathTest struct {
	FS   fs.FS
	Path string
	Want error
}

var installPathTests = map[string]installPathTest{
	"path to readable dir is good": {
		FS: fstest.MapFS{
			"path/to/readable/dir": directory(0444),
		},
		Path: "path/to/readable/dir",
		Want: nil,
	},
	"path to unreadable dir is invalid": {
		FS: fstest.MapFS{
			"path/to/unreadable/dir": directory(0333),
		},
		Path: "path/to/unreadable/dir",
		Want: ErrCannotRead,
	},
	"path to nowhere is invalid": {
		FS: fstest.MapFS{
			"path/to/nowhere": nil,
		},
		Path: "path/to/nowhere",
		Want: ErrNotExist,
	},
	"path to link is invalid": {
		FS: fstest.MapFS{
			"path/to/link": linkTo("some/place"),
		},
		Path: "path/to/link",
		Want: ErrNotDir,
	},
	"path to file is invalid": {
		FS: fstest.MapFS{
			"path/to/file": regularFile(),
		},
		Path: "path/to/file",
		Want: ErrNotDir,
	},
	"path to farm dir is invalid": {
		FS: fstest.MapFS{
			"path/to/farm-dir":       directory(0755),
			"path/to/farm-dir/.farm": regularFile(),
		},
		Path: "path/to/farm-dir",
		Want: ErrIsFarmDir,
	},
	"path to dir inside farm is invalid": {
		FS: fstest.MapFS{
			"path/to/farm-dir/.farm":                   regularFile(),
			"path/to/farm-dir/dir/dir/dir-inside-farm": directory(0755),
		},
		Path: "path/to/farm-dir/dir/dir/dir-inside-farm",
		Want: ErrIsFarmDir,
	},
}

func TestCheckInstallPath(t *testing.T) {
	for name, test := range installPathTests {
		t.Run(name, func(t *testing.T) {
			got := CheckInstallPath(test.FS, test.Path)
			if !errors.Is(got, test.Want) {
				t.Errorf("got error %v, want %v", got, test.Want)
			}
		})
	}
}
