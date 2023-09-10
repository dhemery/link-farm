package cmd

import "flag"

var Relink = &Command{
	Name:      "relink",
	Run:       nil,
	UsageLine: "relink [options] pkg...",
	ShortHelp: "Relink packages",
	FullHelp:  "relink packages",
	Flags:     flag.NewFlagSet("relink", flag.ExitOnError),
}
