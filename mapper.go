package main

import (
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
	return CreateLinkAction{From: targetPath, To: sourcePath}, nil
}

type FileExistsError struct {
	Path string
}

func (e FileExistsError) Error() string {
	return fmt.Sprint("existing file:", e.Path)
}

type NoSuchDirectoryError struct {
	Dir string
}

func (e NoSuchDirectoryError) Error() string {
	return fmt.Sprint("no such directory:", e.Dir)
}
