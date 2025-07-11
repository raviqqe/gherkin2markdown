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

func Run(ss []string, w io.Writer) error {
	args, err := GetArguments(ss)

	if err != nil {
		return err
	} else if args.File == "" {
		return convert.ConvertFiles(args.SrcDir, args.DestDir)
	}

	return convert.ConvertFile(args.File, w)
}
