package main

type CreateLink struct {
	From string
	To   string
}

func (a CreateLink) Perform() error {
	return nil
}
