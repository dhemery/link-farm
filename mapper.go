package main

import (
	"errors"
	"io/fs"
	"path"
	"strings"
)

type Action interface {
	Perform() error
}

type Mapper struct {
	FS        fs.StatFS
	Linker    Symlinker
	SourceDir string
	TargetDir string
}

func (m *Mapper) Map(sourcePath string) (Action, error) {
	relPath := strings.TrimPrefix(sourcePath, m.SourceDir)
	targetPath := path.Join(m.TargetDir, relPath)
	targetEntry, err := m.FS.Stat(targetPath)
	if errors.Is(err, fs.ErrNotExist) {
		return CreateLink{
			Linker: m.Linker,
			From:   targetPath,
			To:     sourcePath,
		}, nil
	}
	sourceEntry, _ := m.FS.Stat(sourcePath)
	if targetEntry.IsDir() && sourceEntry.IsDir() {
		return Descend{}, nil
	}
	return nil, fs.ErrExist
}
