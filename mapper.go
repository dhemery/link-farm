package main

import (
	"errors"
	"io/fs"
	"path"
	"strings"
)

type Action interface {
	Perform() error
}

type Result struct {
	Action  Action
	Descend bool
	Error   error
}

type Mapper struct {
	FS         fs.FS
	Linker     Symlinker
	PackageDir string
	InstallDir string
}

func (m *Mapper) Map(packagePath string) (Action, error) {
	relPath := strings.TrimPrefix(packagePath, m.PackageDir)
	imagePath := path.Join(m.InstallDir, relPath)
	imageItem, err := fs.Stat(m.FS, imagePath)
	packageItem, _ := fs.Stat(m.FS, packagePath)
	if errors.Is(err, fs.ErrNotExist) {
		err = nil
		if packageItem.IsDir() {
			err = fs.SkipDir
		}
		return CreateLink{
			Linker:      m.Linker,
			ImagePath:   imagePath,
			PackagePath: packagePath,
		}, err
	}
	if imageItem.IsDir() && packageItem.IsDir() {
		return nil, nil
	}
	return nil, fs.ErrExist
}
