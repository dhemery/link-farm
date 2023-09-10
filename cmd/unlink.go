package cmd

import "flag"

var Unlink = &Command{
	Name:      "unlink",
	Run:       nil,
	UsageLine: "unlink [options] pkg...",
	ShortHelp: "Unlink packages",
	FullHelp:  "unlink packages",
	Flags:     flag.NewFlagSet("unlink", flag.ExitOnError),
}
