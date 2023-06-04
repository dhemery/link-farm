package main

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"
	"time"
)

type mapperTest struct {
	PackageEntry *fstest.MapFile
	InstallEntry *fstest.MapFile
	WantLink     bool
	WantError    error
}

var mapperTests = map[string]mapperTest{
	"link to package file if install dir has no such path": {
		PackageEntry: regularFile(),
		InstallEntry: nil,
		WantLink:     true,
		WantError:    nil,
	},
	"link to package dir if install dir has no such path": {
		PackageEntry: directory(0644),
		InstallEntry: nil,
		WantLink:     true,
		WantError:    nil,
	},
	"cannot link existing image file to package file": {
		PackageEntry: regularFile(),
		InstallEntry: regularFile(),
		WantLink:     false,
		WantError:    fs.ErrExist,
	},
	"cannot link existing image file to package dir": {
		PackageEntry: directory(0644),
		InstallEntry: regularFile(),
		WantLink:     false,
		WantError:    fs.ErrExist,
	},
	"cannot link existing image dir to package file": {
		PackageEntry: regularFile(),
		InstallEntry: directory(0644),
		WantLink:     false,
		WantError:    fs.ErrExist,
	},
	"continue walking with no action if package and image are both dirs": {
		PackageEntry: directory(0644),
		InstallEntry: directory(0644),
		WantLink:     false, // No link ..
		WantError:    nil,   // ... but no error, so continue walking
	},
}

func TestMapper(t *testing.T) {
	for name, test := range mapperTests {
		t.Run(name, func(t *testing.T) {
			entryPath := "path/to/entry"
			fsys := fstest.MapFS{
				entryPath: test.InstallEntry,
			}
			mapper := Mapper{InstallFS: fsys}

			packageEntry := dirEntry{"entry", test.PackageEntry}
			link, err := mapper.Map(packageEntry, entryPath)
			if link != test.WantLink {
				t.Errorf("got link %#v, want %#v", link, test.WantLink)
			}
			if !errors.Is(err, test.WantError) {
				t.Errorf(`got error %v, want "%v"`, err, test.WantError)
			}
		})
	}
}


func directory(mode fs.FileMode) *fstest.MapFile {
	return &fstest.MapFile{Mode: mode | fs.ModeDir}
}

func regularFile() *fstest.MapFile {
	return &fstest.MapFile{Mode: 0644}
}

func linkTo(p string) *fstest.MapFile {
	return &fstest.MapFile{
		Mode: 0644 | fs.ModeSymlink,
		Data: []byte(p),
	}
}

type dirEntry struct {
	name string
	file *fstest.MapFile
}

func (e dirEntry) Info() (fs.FileInfo, error) {
	return e, nil
}

func (e dirEntry) Mode() fs.FileMode {
	return e.file.Mode
}

func (e dirEntry) IsDir() bool {
	return e.file.Mode.IsDir()
}

func (dirEntry) ModTime() time.Time {
	panic("unimplemented")
}

func (e dirEntry) Name() string {
	return e.name
}

func (e dirEntry) Size() int64 {
	return int64(len(e.file.Data))
}

func (dirEntry) Sys() any {
	return nil
}

func (e dirEntry) Type() fs.FileMode {
	return e.Mode()
}
