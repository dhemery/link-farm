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

type Mapper struct {
	FS         fs.StatFS
	Linker     Symlinker
	PackageDir string
	InstallDir string
}

func (m *Mapper) Map(packagePath string) (Action, error) {
	relPath := strings.TrimPrefix(packagePath, m.PackageDir)
	imagePath := path.Join(m.InstallDir, relPath)
	imageItem, err := m.FS.Stat(imagePath)
	if errors.Is(err, fs.ErrNotExist) {
		return CreateLink{
			Linker:      m.Linker,
			ImagePath:   imagePath,
			PackagePath: packagePath,
		}, nil
	}
	packageItem, _ := m.FS.Stat(packagePath)
	if imageItem.IsDir() && packageItem.IsDir() {
		return Descend{}, nil
	}
	return nil, fs.ErrExist
}
