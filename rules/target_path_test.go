package rules

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"
)

type targetPathTest struct {
	FS          fstest.MapFS
	Path        string
	WantCanLink bool
	WantError   error
}

var targetPathTests = map[string]targetPathTest{
	"path to nowhere can link": {
		FS: fstest.MapFS{
			"path/to/nowhere": nil,
		},
		Path:        "path/to/nowhere",
		WantCanLink: true,
		WantError:   nil,
	},
	"path to file is err exist": {
		FS: fstest.MapFS{
			"path/to/file": regularFile(),
		},
		Path:        "path/to/file",
		WantCanLink: false,
		WantError:   fs.ErrExist,
	},
}

func TestCheckTargetPath(t *testing.T) {
	for name, test := range targetPathTests {
		t.Run(name, func(t *testing.T) {
			canLink, err := CheckTargetPath(test.FS, test.Path)
			if canLink != test.WantCanLink {
				t.Errorf("got can link %t, want %t", canLink, test.WantCanLink)
			}
			if !errors.Is(err, test.WantError) {
				t.Errorf("got error %v, want %v", err, test.WantError)
			}

		})
	}

}
