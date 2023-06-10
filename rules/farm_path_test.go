package rules

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"
)

type farmPathTest struct {
	FS   fs.FS
	Path string
	Want error
}

var farmPathTests = map[string]farmPathTest{
	"path to dir with .farm file is good": {
		FS: fstest.MapFS{
			"path/to/dir-with-farm-file/.farm": regularFile(),
		},
		Path: "path/to/dir-with-farm-file",
		Want: nil,
	},
	"path to dir with no .farm file is invalid": {
		FS: fstest.MapFS{
			"path/to/dir-with-no-farm-file": directory(0755),
		},
		Path: "path/to/dir-with-no-farm-file",
		Want: fs.ErrInvalid,
	},
	"path to file is invalid": {
		FS: fstest.MapFS{
			"path/to/file": regularFile(),
		},
		Path: "path/to/file",
		Want: fs.ErrInvalid,
	},
	"path to link is invalid": {
		FS: fstest.MapFS{
			"path/to/link": linkTo("some/place"),
		},
		Path: "path/to/link",
		Want: fs.ErrInvalid,
	},
	"path to non-existent entry is invalid": {
		FS:   fstest.MapFS{},
		Path: "path/to/nothing",
		Want: fs.ErrInvalid,
	},
}

func TestCheckFarmPath(t *testing.T) {
	for name, test := range farmPathTests {
		t.Run(name, func(t *testing.T) {
			got := CheckIsFarm(test.FS, test.Path)
			if !errors.Is(got, test.Want) {
				t.Errorf("got %v, want %v", got, test.Want)
			}
		})
	}
}
