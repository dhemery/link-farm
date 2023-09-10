package link

import (
	"dhemery.com/duffel/cmd/base"
)

var CmdUnlink = &base.Command{
	Name:            "unlink",
	Run:             runUnlink,
	ArgList:         "pkg...",
	Summary:         "Remove links from target dir to packages in source dir",
	FullDescription: "unlink full description",
	Flags:           linkFlags,
}

func runUnlink(cmd *base.Command, args []string) error {
	return nil
}
