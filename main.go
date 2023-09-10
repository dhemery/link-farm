package main

import (
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
	duffelPath := "duffel"
	installPath := "install"
	packages := []string{
		"shared-1",
		"shared-2",
		"distinct",
	}
	plan.MakePlan(duffelPath, installPath, packages...)
}
