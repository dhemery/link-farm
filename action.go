package main

type source interface{}
type target interface{}

type CreateLink struct {
	Path string
}

func (a CreateLink) Perform( s source, t target) error {
	return nil
}

type Descend struct {
}

func (a Descend) Perform(s source, t target) error {
	return nil
}
