package main

import (
	"fmt"
	"io"
	"os"

	"github.com/raviqqe/gherkin2markdown/convert"
	"github.com/spf13/pflag"
)

const version = "0.1.0"

func main() {
	if err := Run(os.Args[1:], os.Stdout); err != nil {
		if _, err := fmt.Fprintln(os.Stderr, err); err != nil {
			panic(err)
		}

		os.Exit(1)
	}
}

// Run executes the CLI command.
func Run(ss []string, w io.Writer) error {
	args, err := GetArguments(ss)

	if err != nil {
		return err
	} else if args.Help {
		pflag.PrintDefaults()
		return nil
	} else if args.Version {
		_, err := fmt.Fprintln(w, version)
		return err
	} else if args.SrcDir == "" || args.DestDir == "" {
		return fmt.Errorf("source and destination directories required")
	}

	return convert.FeatureFiles(args.SrcDir, args.DestDir)
}
