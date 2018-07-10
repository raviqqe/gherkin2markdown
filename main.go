package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if err := command(os.Args[1:], os.Stdout); err != nil {
		if _, err := fmt.Fprintln(os.Stderr, err); err != nil {
			panic(err)
		}

		os.Exit(1)
	}
}

func command(ss []string, w io.Writer) error {
	args := getArguments(ss)

	if args.File == "" {
		return convertFiles(args.SrcDir, args.DestDir)
	}

	return convertFile(args.File, w)
}
