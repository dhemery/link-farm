package link

import (
	"flag"

	"dhemery.com/duffel/cmd/base"
)

const linkDescription = `
duffel link creates links in the target directory that point to
corresponding entries in the file trees of the named packages.

The default target directory is the parent of the current
working directory. To specify a different target directory, use
the -target option.

duffel looks for the named packages in the source directory. The
default source directory is the current working directory. To
specify a different source directory, use the -source option.

duffel link evaluates all planned actions before performing any.
If any planned action is invalid, duffel link prints an error
message and exits without performing any actions. Use the -plan
option to preview the plan.
`

var (
	CmdLink = &base.Command{
		Name:            "link",
		Run:             runLink,
		ArgList:         "pkg...",
		Summary:         "Create links to packages",
		FullDescription: linkDescription,
		Flags:           linkFlags,
	}

	linkFlags = flag.NewFlagSet("", flag.ExitOnError)
	onlyPlan  *bool
	sourceDir *string
	targetDir *string
	verbose   *bool
)

func init() {
	onlyPlan = linkFlags.Bool("plan", true, "print the planned actions without executing them")
	sourceDir = linkFlags.String("source", ".", "set source directory to `dir`")
	targetDir = linkFlags.String("target", "..", "set target directory to `dir`")
	verbose = linkFlags.Bool("verbose", false, "print each action immediately before executing it")
}

func runLink(cmd *base.Command, args []string) error {
	return nil
}
