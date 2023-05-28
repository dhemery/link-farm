package main

type CreateLinkAction struct {
	From string
	To   string
}

func (a CreateLinkAction) Perform() error {
	return nil
}
