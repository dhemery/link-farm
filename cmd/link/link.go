package link

import (
	"flag"

	"dhemery.com/duffel/cmd/base"
)

var (
	linkFlags = flag.NewFlagSet("", flag.ExitOnError)
	CmdLink   = &base.Command{
		Name:            "link",
		Run:             runLink,
		ArgList:         "pkg...",
		Summary:         "Create links from target dir to packages in source dir",
		FullDescription: "link full description",
		Flags:           linkFlags,
	}
	onlyPlan  *bool
	sourceDir *string
	targetDir *string
)

func init() {
	onlyPlan = linkFlags.Bool("plan", true, "print plan without running actions")
	sourceDir = linkFlags.String("source", ".", "set source directory to `dir`")
	targetDir = linkFlags.String("target", "..", "set target directory to `dir`")
}

func runLink(cmd *base.Command, args []string) error {
	return nil
}
