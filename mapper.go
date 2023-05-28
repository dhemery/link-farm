package main

import (
	"errors"
	"fmt"
	"io/fs"
	"path"
	"strings"
)

type Action interface {
	Perform() error
}

type Mapper struct {
	FS        fs.StatFS
	SourceDir string
	TargetDir string
}

func (m *Mapper) Map(sourcePath string) (Action, error) {
	relPath := strings.TrimPrefix(sourcePath, m.SourceDir)
	targetPath := path.Join(m.TargetDir, relPath)

	_, err := m.FS.Stat(targetPath)
	if errors.Is(err, fs.ErrNotExist) {
		return CreateLinkAction{From: targetPath, To: sourcePath}, nil
	}
	return nil, fmt.Errorf("%s: %w", targetPath, fs.ErrExist)
}

type NoSuchDirectoryError struct {
	Dir string
}

func (e NoSuchDirectoryError) Error() string {
	return fmt.Sprint("no such directory:", e.Dir)
}
