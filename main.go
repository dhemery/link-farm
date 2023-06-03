package main


import (
	"fmt"
	"os"
	"path"
)


func main() {
	farm := "example/farm"
	target := "example/target"
	if err :=linkFarm(farm, target); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func linkFarm (farm, target string) error {
	errs := []error{}
	farmEntries, err := os.ReadDir(farm)
	if err != nil {
		return err
	}
	for _, e := range farmEntries {
		if !e.IsDir() {
			continue
		}
		pkg := path.Join(farm, e.Name())
		if err := linkPackage(pkg, target); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("link farm errprs: %s", errs)
	}
	return nil
}

func linkPackage(pkg, target string) error {
	fmt.Println("linking package", pkg, "to target", target)
	return nil
}

