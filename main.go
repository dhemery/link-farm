package main

import (
	"fmt"
	"io/fs"
	"os"
	"path"
)


func main() {
	fsys := os.DirFS("example")
	if err :=linkFarm(fsys, "farm", "target"); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func linkFarm (fsys fs.FS, farm, target string) error {
	errs := []error{}
	farmEntries, err := fs.ReadDir(fsys, farm)
	if err != nil {
		return err
	}
	for _, e := range farmEntries {
		if !e.IsDir() {
			continue
		}
		pkg := path.Join(farm, e.Name())
		if err := linkPackage(fsys, pkg, target); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("link farm errprs: %s", errs)
	}
	return nil
}

func linkPackage(fsys fs.FS, pkg, target string) error {
	m := Mapper{FS: fsys, PackageDir: pkg, InstallDir: target}
	return fs.WalkDir(fsys, pkg, entryAction(m))
}

func entryAction(m Mapper) fs.WalkDirFunc {
	return func(path string, d os.DirEntry, errIn error) error {
		mapped, err := m.Map(path)
		fmt.Printf("mapping %s with %#v\n", path,  mapped)
		return err
	}
}
