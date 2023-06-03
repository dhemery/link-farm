package main

import (
	"fmt"
	"io/fs"
	"os"
	"path"
)

func main() {
	fsys := os.DirFS("example")
	if err := linkFarm(fsys, "farm", "target"); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func linkFarm(fsys fs.FS, farm, target string) error {
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
	installFS, _ := fs.Sub(fsys, target)
	m := Mapper{InstallFS: installFS}
	return fs.WalkDir(fsys, pkg, entryAction(fsys, m))
}

func entryAction(fsys fs.FS, m Mapper) fs.WalkDirFunc {
	return func(path string, d os.DirEntry, errIn error) error {
		packageEntry, _ := fs.Stat(fsys, path)
		link, err := m.Map(packageEntry, path)
		fmt.Println("Mapped", path, "link", link, "err", err)
		if link || err != nil {
			return fs.SkipDir
		}
		return nil
	}
}
