package main

import (
	"fmt"
	"io"
	"os"

	"github.com/raviqqe/gherkin2markdown/convert"
)

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
	} else if args.File == "" {
		return convert.FeatureFiles(args.SrcDir, args.DestDir)
	}

	return convert.FeatureFile(args.File, w)
}
