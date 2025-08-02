package main

import (
	"os"

	"github.com/spf13/pflag"
)

type Arguments struct {
	SrcDir  string
	DestDir string
	Help    bool
	Version bool
}

func GetArguments(ss []string) (Arguments, error) {
	s := pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError)
	args := Arguments{}

	s.BoolVar(&args.Help, "help", false, "show help")
	s.BoolVar(&args.Version, "version", false, "show version")

	if err := s.Parse(ss); err != nil {
		return args, err
	}

	if ds := s.Args(); len(ds) == 2 {
		args.SrcDir = ds[0]
		args.DestDir = ds[1]
	}

	return args, nil
}
