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
	ErrNotDir        = errors.New("is not a directory")
	ErrNotRegular    = errors.New("is not a regular file")
	ErrNotExist      = errors.New("does not exist")
	ErrCannotRead    = errors.New("cannot read")
	ErrCannotWrite   = errors.New("cannot write")
	ErrCannotExecute = errors.New("cannot execute")
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
	return checkCanWrite(info)
}

func checkReadableDir(info fs.FileInfo) error {
	if !info.IsDir() {
		return ErrNotDir
	}
	return checkCanRead(info)
}

const (
	readBits  = 0444
	writeBits = 0222
	execBits  = 0111
)

func checkCanRead(info fs.FileInfo) error {
	perm := info.Mode().Perm()
	if perm&readBits == 0 {
		return fmt.Errorf("%w: perm %04o", ErrCannotRead, perm)
	}
	return nil
}

func checkCanWrite(info fs.FileInfo) error {
	perm := info.Mode().Perm()
	if perm&writeBits == 0 {
		return fmt.Errorf("%w: perm %04o", ErrCannotWrite, perm)
	}
	return nil
}

func checkCanExecute(info fs.FileInfo) error {
	perm := info.Mode().Perm()
	if perm&execBits == 0 {
		return fmt.Errorf("%w: perm %04o", ErrCannotExecute, perm)
	}
	return nil
}
