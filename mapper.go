package main

import (
	"errors"
	"io/fs"
)

type Mapper struct {
	InstallFS fs.FS
}

func (m *Mapper) Map(packageEntry fs.FileInfo, installPath string) (bool, error) {
	installEntry, err := fs.Stat(m.InstallFS, installPath)
	if errors.Is(err, fs.ErrNotExist) {
		return true, nil
	}
	if installEntry.IsDir() && packageEntry.IsDir() {
		return false, nil
	}
	return false, fs.ErrExist
}
