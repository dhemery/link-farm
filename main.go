package main

import (
	"flag"
	"fmt"
	"os"

	"dhemery.com/duffel/cmd/base"
	"dhemery.com/duffel/cmd/help"
	"dhemery.com/duffel/cmd/link"
)

func init() {
	base.Commands = []*base.Command{
		link.CmdLink,
		link.CmdUnlink,
		help.CmdHelp,
	}
}

func main() {
	flag.Usage = help.Usage
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		flag.Usage()
		os.Exit(2)
	}

	cmdName := args[0]
	cmd, ok := base.FindCommand(cmdName)
	if !ok {
		fmt.Fprintln(os.Stderr, "no such command:", cmdName)
		flag.Usage()
		os.Exit(2)
	}

	err := cmd.Run(cmd, args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
