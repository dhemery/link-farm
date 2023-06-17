package rules

import (
	"errors"
	"fmt"
	"io/fs"
	"path"
)

const (
	fsRoot = "."
)

var (
	ErrNotDir     = errors.New("is not a directory")
	ErrNotRegular = errors.New("is not a regular file")
	ErrNotExist   = errors.New("does not exist")
)

func checkCanCreate(f fs.FS, p string) error {
	parent := path.Dir(p)
	info, err := fs.Stat(f, parent)
	if errors.Is(err, fs.ErrNotExist) {
		return checkCanCreate(f, parent)
	}
	if err != nil {
		return err
	}
	mode := info.Mode()
	if !isWriteable(mode) {
		return fmt.Errorf("%s: cannot write (mode %o): %w", p, mode, fs.ErrPermission)
	}
	return nil
}

func checkReadableDir(info fs.FileInfo) error {
	if !info.IsDir() {
		return fmt.Errorf("not a directory: %w", fs.ErrInvalid)
	}
	mode := info.Mode()
	if !isReadable(mode) {
		return fmt.Errorf("cannot read (mode %o): %w", mode, fs.ErrPermission)
	}
	return nil
}

func isReadable(m fs.FileMode) bool {
	const readBits = 0444
	return m&readBits != 0
}

func isWriteable(m fs.FileMode) bool {
	const writeBits = 0222
	return m&writeBits != 0
}
