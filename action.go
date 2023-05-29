package main

import "errors"

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

type Descend struct{}

func (a Descend) Perform() error {
	return errors.New("Descend action was performed")
}
