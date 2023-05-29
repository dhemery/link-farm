package main

type CreateLinkAction struct {
	From string
	To   string
}

func (a CreateLinkAction) Perform() error {
	return nil
}

type ReplaceWithLink struct {
	From string
	To   string
}

func (a ReplaceWithLink) Perform() error {
	return nil
}
