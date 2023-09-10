package cmd

import "flag"

var Link = &Command{
	Name:      "link",
	Run:       nil,
	UsageLine: "link [options] pkg...",
	ShortHelp: "Link packages",
	FullHelp:  "link packages",
	Flags:     flag.NewFlagSet("link", flag.ExitOnError),
}
