package main

import (
	"github.com/docopt/docopt-go"
)

const usage = `Gherkin to Markdown converter

Usage:
	gherkin2markdown <file>
	gherkin2markdown <src_dir> <dest_dir>

Options:
	-h, --help  Show this help.`

// Arguments is the available CLI arguments.
type Arguments struct {
	File    string `docopt:"<file>"`
	SrcDir  string `docopt:"<src_dir>"`
	DestDir string `docopt:"<dest_dir>"`
}

// GetArguments return the CLI arguments.
func GetArguments(ss []string) (arguments, error) {
	args := arguments{}
	err := parseArguments(usage, ss, &args)
	return args, err
}

func parseArguments(u string, ss []string, args interface{}) error {
	opts, err := docopt.ParseArgs(u, ss, "")

	if err != nil {
		return err
	}

	return opts.Bind(args)
}
