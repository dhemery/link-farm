package main

import (
	"errors"
	"fmt"
	"io/fs"
	"path"
)

func CheckPackagePath(f fs.FS, p string) error {
	info, err := fs.Stat(f, p)
	if err != nil {
		return err
	}

	fmt.Printf("check readdable %s mode %o\n", p, info.Mode())
	return checkReadableDir(info)
}

func CheckInstallPath(f fs.FS, p string) error {
	info, err := fs.Stat(f, p)
	if err != nil {
		return err
	}
	return checkReadableDir(info)
}

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

const (
	readBits = 0444
	writeBits = 0222
)

func isReadable(m fs.FileMode) bool {
	return m&readBits != 0
}

func isWriteable(m fs.FileMode) bool {
	return m&writeBits != 0
}
