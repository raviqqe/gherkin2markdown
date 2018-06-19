package main

import (
	"github.com/docopt/docopt-go"
)

var usage = `Gherkin to Markdown converter

Usage:
	gherkin2markdown <file>
	gherkin2markdown <srcdir> <destdir>

Options:
	-h, --help  Show this help.`

type arguments struct {
	File    string `docopt:"<file>"`
	SrcDir  string `docopt:"<srcdir>"`
	DestDir string `docopt:"<destdir>"`
}

func getArguments(ss []string) arguments {
	args := arguments{}
	parseArguments(usage, ss, &args)
	return args
}

func parseArguments(u string, ss []string, args interface{}) {
	opts, err := docopt.ParseArgs(u, ss, "0.1.0")

	if err != nil {
		panic(err)
	}

	if err := opts.Bind(args); err != nil {
		panic(err)
	}
}
