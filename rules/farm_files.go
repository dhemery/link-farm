package rules

import (
	"fmt"
	"io/fs"
	"path"
)

const (
	farmFileName = ".farm"
)

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
