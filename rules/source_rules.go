package rules

import (
	"fmt"
	"io/fs"
)

func CheckSourceRules(f fs.FS, path string) error {
	info, err := fs.Stat(f, path)
	if err != nil {
		return fmt.Errorf("source path: %w", err)
	}

	mode := info.Mode()
	if !mode.IsDir() && !mode.IsRegular() {
		return fmt.Errorf("source path: %w", fs.ErrInvalid)
	}
	return nil
}
