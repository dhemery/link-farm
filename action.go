package main

import (
	"errors"
)

type Symlinker interface {
	Symlink(oldname, newname string) error
}

type CreateLink struct {
	Linker      Symlinker
	ImagePath   string
	PackagePath string
}

func (a CreateLink) Perform() error {
	return nil
}

type MapChildren struct{}

func (a MapChildren) Perform() error {
	return errors.New("MapChildren action was performed")
}

