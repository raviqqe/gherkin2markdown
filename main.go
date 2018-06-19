package main

import (
	"fmt"
	"os"

	"github.com/cucumber/gherkin-go"
)

func main() {
	err := command()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func command() error {
	f, err := os.Open(os.Args[1])

	if err != nil {
		return err
	}

	d, err := gherkin.ParseGherkinDocument(f)

	if err != nil {
		return err
	}

	fmt.Print(newRenderer().Render(d))

	return nil
}
