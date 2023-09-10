package cmd

import "flag"

var All = []*Command{}

type Command struct {
	Name      string
	Run       func(cmd *Command, args []string) error
	UsageLine string
	ShortHelp string
	FullHelp  string
	Flags     *flag.FlagSet
}
