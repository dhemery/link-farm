package main

import (
	"flag"
	"fmt"
	"os"

	"dhemery.com/duffel/cmd"
	"dhemery.com/duffel/plan"
)

func init() {
	cmd.All = []*cmd.Command{
		cmd.Link,
		cmd.Relink,
		cmd.Unlink,
	}
}

func main() {
	flag.Usage = cmd.Duffel.Usage
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		flag.Usage()
		os.Exit(2)
	}

	duffelPath := "duffel"
	installPath := "install"
	packages := []string{
		"shared-1",
		"shared-2",
		"distinct",
	}

	c,ok := cmd.FindCommand(args[0])
	if !ok {
		flag.Usage()
		os.Exit(2)
	}

	err := c.Run(c, args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	err = plan.MakePlan(duffelPath, installPath, packages...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
