package main

import (
	"log"
	"os"
	"path"

	"dhemery.com/duffel/rules"
)

func main() {
	root := os.DirFS("example")
	duffelPath := "duffel"
	installPath := "install"
	packages := []string{
		"shared-1",
		"shared-2",
		"distinct",
	}

	exitCode := 0

	if err := rules.CheckIsDuffelDir(root, duffelPath); err != nil {
		log.Printf("invalid duffel path %s: %s", installPath, err)
		exitCode = 1
	}
	if err := rules.CheckInstallPath(root, installPath); err != nil {
		log.Printf("invalid install path %s: %s", installPath, err)
		exitCode = 1
	}

	for _, name := range packages {
		packagePath := path.Join(duffelPath, name)
		if err := rules.CheckPackagePath(root, packagePath); err != nil {
			log.Printf("invalid package path %s: %s", packagePath, err)
			exitCode = 1
		}
	}

	os.Exit(exitCode)
}
