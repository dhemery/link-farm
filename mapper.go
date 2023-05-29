package main

import (
	"errors"
	"fmt"
	"io/fs"
)

type Action interface {
	Perform(s source, t target) error
}

type Mapper struct {
	Source fs.StatFS
	Target fs.StatFS
}

func (m *Mapper) Map(path string) (Action, error) {
	targetEntry, err := m.Target.Stat(path)
	if errors.Is(err, fs.ErrNotExist) {
		return CreateLink{Path: path}, nil
	}
	if targetEntry.IsDir() {
		entries, _ := fs.ReadDir(m.Target, path)
		if len(entries) == 0 {
			return ReplaceWithLink{Path: path}, nil
		}

	}
	sourceEntry, _ := m.Source.Stat(path)
	if sourceEntry.IsDir() {
		return nil, fmt.Errorf("%s: %w", path, fs.ErrExist)
	} else {
		return nil, fs.ErrExist
	}
}
