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
	Mapper      Mapper
	PackagePath string
	Action      Action
	Error       error
}

var mapperTests = map[string]mapperTest{
	"create link to package file if install dir has no such path": {
		Mapper: Mapper{
			PackageDir: "package-dir",
			InstallDir: "install-dir",
			FS: fstest.MapFS{
				"package-dir/path/to/item": regularFile(),
				// Install dir has no such path
			},
			Linker: testLinker{},
		},
		PackagePath: "package-dir/path/to/item",
		Action: CreateLink{
			Linker:      testLinker{},
			ImagePath:   "install-dir/path/to/item",
			PackagePath: "package-dir/path/to/item",
		},
	},
	"create link to package dir if install dir has no such path": {
		Mapper: Mapper{
			PackageDir: "package-dir",
			InstallDir: "install-dir",
			FS: fstest.MapFS{
				"package-dir/path/to/item": directory(),
				// Install dir has no such path
			},
			Linker: testLinker{},
		},
		PackagePath: "package-dir/path/to/item",
		Action: CreateLink{
			Linker:      testLinker{},
			ImagePath:   "install-dir/path/to/item",
			PackagePath: "package-dir/path/to/item",
		},
	},
	"cannot link existing image file to package file": {
		Mapper: Mapper{
			PackageDir: "package-dir",
			InstallDir: "install-dir",
			FS: fstest.MapFS{
				"package-dir/path/to/item": regularFile(),
				"install-dir/path/to/item": regularFile(),
			},
		},
		PackagePath: "package-dir/path/to/item",
		Error:       fs.ErrExist,
	},
	"cannot link existing image file to package dir": {
		Mapper: Mapper{
			PackageDir: "package-dir",
			InstallDir: "install-dir",
			FS: fstest.MapFS{
				"package-dir/path/to/item": directory(),
				"install-dir/path/to/item": regularFile(),
			},
		},
		PackagePath: "package-dir/path/to/item",
		Error:       fs.ErrExist,
	},
	"cannot link existing image dir to package file": {
		Mapper: Mapper{
			PackageDir: "package-dir",
			InstallDir: "install-dir",
			FS: fstest.MapFS{
				"package-dir/path/to/item": regularFile(),
				"install-dir/path/to/item": directory(),
			},
		},
		PackagePath: "package-dir/path/to/item",
		Error:       fs.ErrExist,
	},
	"descend if package and image are both dirs": {
		Mapper: Mapper{
			PackageDir: "package-dir",
			InstallDir: "install-dir",
			FS: fstest.MapFS{
				"package-dir/path/to/item": directory(),
				"install-dir/path/to/item": directory(),
			},
		},
		PackagePath: "package-dir/path/to/item",
		Action:      Descend{},
	},
}

func TestMapper(t *testing.T) {
	for name, test := range mapperTests {
		t.Run(name, func(t *testing.T) {
			action, err := test.Mapper.Map(test.PackagePath)
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
