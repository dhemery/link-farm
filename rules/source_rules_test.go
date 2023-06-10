package rules

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"
)

type sourcePathRulesTest struct {
	FS   fstest.MapFS
	Path string
	Want error
}

var sourcePathRulesTests = map[string]sourcePathRulesTest{
	"file is good": {
		FS: fstest.MapFS{
			"path/to/file": regularFile(),
		},
		Path: "path/to/file",
		Want: nil,
	},
	"dir is good": {
		FS: fstest.MapFS{
			"path/to/dir": directory(0x755),
		},
		Path: "path/to/dir",
		Want: nil,
	},
	"link is invalid": {
		FS: fstest.MapFS{
			"path/to/link": linkTo("some/place"),
		},
		Path: "path/to/link",
		Want: fs.ErrInvalid,
	},
}

func TestSourcePathRules(t *testing.T) {
	for name, test := range sourcePathRulesTests {
		t.Run(name, func(t *testing.T) {
			got := CheckSourcePathRules(test.FS, test.Path)
			if !errors.Is(got, test.Want) {
				t.Errorf("got error %v, want %v", got, test.Want)
			}
		})
	}
}
