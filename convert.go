package main

import (
	"fmt"
	"io"
	"os"

	"github.com/cucumber/gherkin-go"
)

func convertFile(s string, w io.Writer) error {
	f, err := os.Open(s)

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

func convertFiles(s, d string) error {
	panic("unimplemented")
}
