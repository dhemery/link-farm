package cmd

import (
	"flag"
	"fmt"
	"os"
)

var All = []*Command{}

type Command struct {
	Name      string
	Run       func(cmd *Command, args []string) error
	UsageLine string
	ShortHelp string
	FullHelp  string
	Flags     *flag.FlagSet
}

func (cmd *Command) Usage() {
	fmt.Fprintln(os.Stderr, cmd.UsageLine)
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, cmd.FullHelp)
}

func FindCommand(name string) (*Command, bool) {
	for _, c := range All {
		if c.Name == name {
			return c, true
		}
	}
	return nil, false
}
