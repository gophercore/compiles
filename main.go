package main

import (
	_ "embed"
	"flag"
	"log"
	"os"

	"golang.org/x/tools/go/packages"
)

const (
	codeError = 1
	codeUsage = 2
)

func main() {
	var excludeTests bool
	var buildTags string
	flag.BoolVar(&excludeTests, "exclude-tests", false, "do not check test-only package or test executables")
	flag.StringVar(&buildTags, "tags", "", "comma-separated list of build tags to include")

	log.SetFlags(0) // reset flags to remove time prefix
	log.SetPrefix("compiles: ")

	flag.Parse()
	if len(flag.Args()) == 0 {
		log.Print("no packages specified")
		os.Exit(codeUsage)
	}

	cfg := &packages.Config{
		BuildFlags: []string{"-tags=" + buildTags},
		Mode:       packages.LoadAllSyntax | packages.NeedModule,
		Tests:      !excludeTests,
	}
	initial, err := packages.Load(cfg, flag.Args()...)
	if err != nil {
		log.Fatalf("failed to load packages: %v", err)
	}
	if len(initial) == 0 {
		log.Fatalf("no packages found")
	}
	if packages.PrintErrors(initial) > 0 {
		os.Exit(codeError)
	}
}
