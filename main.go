package main

import (
	"fmt"
	"io"
	"os"

	"github.com/cucumber/gherkin-go"
)

func main() {
	err := command(os.Args[1:], os.Stdout)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func command(ss []string, w io.Writer) error {
	f, err := os.Open(ss[0])

	if err != nil {
		return err
	}

	d, err := gherkin.ParseGherkinDocument(f)

	if err != nil {
		return err
	}

	fmt.Fprint(w, newRenderer().Render(d))

	return nil
}
