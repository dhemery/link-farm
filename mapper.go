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

	target, err := m.FS.Stat(targetPath)
	if errors.Is(err, fs.ErrNotExist) {
		return CreateLinkAction{From: targetPath, To: sourcePath}, nil
	}
	if target.IsDir() {
		entries, _ := fs.ReadDir(m.FS, targetPath)
		if len(entries) == 0 {
			return ReplaceWithLink{From: targetPath, To: sourcePath}, nil
		}

	}
	source, _ := m.FS.Stat(sourcePath)
	if source.IsDir() {
		return m.mapDir(source, sourcePath, target, targetPath)
	} else {
		return nil, fs.ErrExist
	}
}

func (m *Mapper) mapDir(source fs.FileInfo, sourcePath string,
	target fs.FileInfo, targetPath string) (Action, error) {
	return nil, fmt.Errorf("%s: %w", targetPath, fs.ErrExist)
}

type NoSuchDirectoryError struct {
	Dir string
}

func (e NoSuchDirectoryError) Error() string {
	return fmt.Sprint("no such directory:", e.Dir)
}
