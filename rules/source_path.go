package rules

import (
	"io/fs"
)

func CheckSourcePath(f fs.FS, path string) error {
	info, err := fs.Stat(f, path)
	if err != nil {
		return ErrNotExist
	}

	mode := info.Mode()
	if !mode.IsDir() && !mode.IsRegular() {
		return ErrNotFileOrDir
	}
	return nil
}
