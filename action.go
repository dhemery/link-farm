package main

type source interface{}
type target interface{}

type CreateLink struct {
	Path string
}

func (a CreateLink) Perform( s source, t target) error {
	return nil
}

type ReplaceWithLink struct {
	Path string
}

func (a ReplaceWithLink) Perform(s source, t target) error {
	return nil
}
