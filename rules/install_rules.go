package rules

import (
	"io/fs"
)

func CheckInstallPath(f fs.FS, p string) error {
	info, err := fs.Stat(f, p)
	if err != nil {
		return err
	}
	if err = checkReadableDir(info); err != nil {
		return err
	}
	return checkNotInFarm(f, p)
}
