package main

import (
	"fmt"
	"io/fs"
	"os"
	"path"
)

type Planner struct {
	FS          fs.FS
	FarmRoot    string
	InstallRoot string
	Packages    []string
}

type Plan struct {
	Links    []string
	Problems []error
}

func (p *Plan) AddProblem(problem error) {
	p.Problems = append(p.Problems, problem)
}

func (p *Plan) AddLink(link string) {
	p.Links = append(p.Links, link)
}

func (p *Plan) Execute() error {
	for _, path := range p.Links {
		fmt.Println("link", path)
	}
	return nil
}

func main() {
	rootFS := os.DirFS("example")
	planner := Planner{
		FS:          rootFS,
		FarmRoot:    "farm",
		InstallRoot: "install",
		Packages: []string{
			"shared-1",
			"shared-2",
			"distinct",
		},
	}
	plan := planner.Plan()
	problems := plan.Problems
	if len(problems) != 0 {
		for _, p := range problems {
			fmt.Fprintln(os.Stderr, p)
		}
		os.Exit(1)
	}
	if err := plan.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func (p *Planner) Plan() *Plan {
	plan := &Plan{}

	installFS, err := fs.Sub(p.FS, p.InstallRoot)
	if err != nil {
		plan.AddProblem(fmt.Errorf("Make %s FS: %w", p.InstallRoot, err))
		return plan
	}

	mapper := Mapper{InstallFS: installFS}
	for _, packageName := range p.Packages {
		p.PlanPackage(packageName, mapper, plan)
	}

	return plan
}

func (p *Planner) PlanPackage(pkgName string, mapper Mapper, plan *Plan) {
	packagePath := path.Join(p.FarmRoot, pkgName)
	packageFS, err := fs.Sub(p.FS, packagePath)
	if err != nil {
		plan.AddProblem(fmt.Errorf("Make package %s FS: %w", packagePath, err))
		return
	}
	packageDirs, err := fs.ReadDir(packageFS, ".")
	if err != nil {
		plan.AddProblem(fmt.Errorf("Read package %s: %w", pkgName, err))
		return
	}
	for _, dir := range packageDirs {
		if err := fs.WalkDir(packageFS, dir.Name(), p.Walker(packageFS, mapper, plan)); err != nil {
			plan.AddProblem(fmt.Errorf("Walk package %s: %w", pkgName, err))
		}
	}
}

func (p *Planner) Walker(fsys fs.FS, mapper Mapper, plan *Plan) fs.WalkDirFunc {
	return func(path string, d os.DirEntry, errIn error) error {
		packageEntry, _ := fs.Stat(fsys, path)
		link, err := mapper.Map(packageEntry, path)
		if link {
			plan.AddLink(path)
		}
		if err != nil {
			plan.AddProblem(err)
		}
		if link || err != nil {
			return fs.SkipDir
		}
		return nil
	}
}
