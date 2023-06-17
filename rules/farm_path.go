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
	ErrNoFarmFile         = errors.New("has no " + farmFileName + " file")
	ErrFarmFileNotRegular = fmt.Errorf("%s %w", farmFileName, ErrNotRegularFile)
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
		return ErrNoFarmFile
	}
	if !farmFileInfo.Mode().IsRegular() {
		return ErrFarmFileNotRegular
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
		return fmt.Errorf("in farm %s: %w", p, fs.ErrPermission)
	}
	return nil
}
