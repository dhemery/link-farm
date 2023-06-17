package rules

import (
	"errors"
	"fmt"
	"io/fs"
	"path"
)

const (
	farmFileName = ".farm"
)

var (
	ErrNotFarmDir = errors.New("is not a farm dir")
	ErrIsFarmDir  = errors.New("is a farm dir")
)

func CheckIsFarm(f fs.FS, p string) error {
	farmInfo, err := fs.Stat(f, p)
	if err != nil {
		return ErrNotExist
	}
	if !farmInfo.IsDir() {
		return ErrNotDir
	}

	farmFilePath := path.Join(p, farmFileName)
	farmFileInfo, err := fs.Stat(f, farmFilePath)
	if err != nil {
		return ErrNotFarmDir
	}
	if !farmFileInfo.Mode().IsRegular() {
		return fmt.Errorf("%s: %w", farmFilePath, ErrNotFile)
	}

	return nil
}

func checkNotInFarm(f fs.FS, p string) error {
	if err := checkNotFarm(f, p); err != nil {
		return err
	}
	if p == fsRoot {
		return nil
	}
	parent := path.Dir(p)
	return checkNotInFarm(f, parent)
}

func checkNotFarm(f fs.FS, p string) error {
	farmFilePath := path.Join(p, farmFileName)
	_, err := fs.Stat(f, farmFilePath)
	if err == nil {
		return fmt.Errorf("%s: %w", p, ErrIsFarmDir)
	}
	return nil
}
