package main

import (
	"errors"
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
	sourceEntry, _ := m.Source.Stat(path)
	if targetEntry.IsDir() && sourceEntry.IsDir() {
		return Descend{}, nil
	}
	return nil, fs.ErrExist
}
