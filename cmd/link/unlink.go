package link

import (
	"dhemery.com/duffel/cmd/base"
)

const unlinkDescription = `
duffel unlink removes links in the target directory that point to
corresponding entries in the file trees of the named packages.
`

var (
	CmdUnlink = &base.Command{
		Name:            "unlink",
		Run:             runUnlink,
		ArgList:         "pkg...",
		Summary:         "Remove links to packages",
		FullDescription: unlinkDescription,
		Flags:           linkFlags,
	}
)

func runUnlink(cmd *base.Command, args []string) error {
	return nil
}
