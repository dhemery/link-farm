package main

import (
	"fmt"
	"path"
	"strings"
)

type Action interface {
	Perform() error
}

type Mapper struct {
	SourceDir string
	TargetDir string
}

func (m *Mapper) Map(sourcePath string) (Action, error) {
	relPath := strings.TrimPrefix(sourcePath, m.SourceDir)
	targetPath := path.Join(m.TargetDir, relPath)
	return CreateLink{From: targetPath, To: sourcePath}, nil
}

type NoSuchDirectory struct {
	Dir string
}

func (e NoSuchDirectory) Error() string {
	return fmt.Sprint("no such directory:", e.Dir)
}
