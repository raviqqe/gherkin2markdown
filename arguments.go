package main

import (
	"github.com/spf13/pflag"
)

type Arguments struct {
	SrcDir  string
	DestDir string
	Help    bool
	Version bool
}

// GetArguments return the CLI arguments.
func GetArguments(ss []string) Arguments {
	args := Arguments{}

	pflag.BoolVar(&args.Help, "help", false, "show help")
	pflag.BoolVar(&args.Version, "version", false, "show version")
	pflag.Parse()

	if ds := pflag.Args(); len(ds) == 2 {
		args.SrcDir = ds[0]
		args.DestDir = ds[1]
	}

	return args
}
