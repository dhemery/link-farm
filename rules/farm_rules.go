package rules

import (
	"fmt"
	"io/fs"
	"path"
)

const (
	farmFileName = ".farm"
)

func CheckIsFarm(f fs.FS, p string) error {
	farmFilePath := path.Join(p, farmFileName)
	info, err := fs.Stat(f, farmFilePath)
	if err != nil {
		return fmt.Errorf("missing %s file: %w", farmFileName, fs.ErrInvalid)
	}
	if !info.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file: %w", farmFilePath, fs.ErrInvalid)
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
