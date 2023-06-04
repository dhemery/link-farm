package rules

import (
	"io/fs"
)

func CheckPackagePath(f fs.FS, p string) error {
	info, err := fs.Stat(f, p)
	if err != nil {
		return err
	}

	return checkReadableDir(info)
}
