package cmd

import "flag"

var Duffel = &Command{
	Name:      "duffel",
	Run:       nil,
	UsageLine: "duffel <command> pkg...",
	ShortHelp: "Manage package installation",
	FullHelp:  "Manage package installation",
	Flags:     flag.NewFlagSet("link", flag.ExitOnError),
}
