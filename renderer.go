package main

import (
	"strings"

	"github.com/cucumber/gherkin-go"
)

type renderer struct {
	*strings.Builder
}

func newRenderer() renderer {
	return renderer{&strings.Builder{}}
}

func (r renderer) Render(d *gherkin.GherkinDocument) string {
	r.render(d.Feature)

	return r.Builder.String()
}

func (r renderer) render(x interface{}) {
	switch x := x.(type) {
	case *gherkin.Feature:
		r.writeLine("# " + x.Name)
		r.writeDescription(x.Description)

		for _, x := range x.Children {
			r.writeLine("")
			r.render(x)
		}
	case *gherkin.Background:
		r.writeLine("## Background (" + x.Name + ")")
		r.writeDescription(x.Description)

		for _, s := range x.Steps {
			r.writeLine("")
			r.render(s)
		}
	case *gherkin.Scenario:
		r.writeLine("## " + x.Name)
		r.writeDescription(x.Description)

		for _, s := range x.Steps {
			r.writeLine("")
			r.render(s)
		}
	case *gherkin.Step:
		r.writeLine("_" + strings.TrimSpace(x.Keyword) + "_ " + x.Text)

		if x.Argument != nil {
			r.writeLine("")
			r.render(x.Argument)
		}
	case *gherkin.DocString:
		r.writeLine("```")
		r.writeLine(x.Content)
		r.writeLine("```")
	default:
		panic("unreachable")
	}
}

func (r renderer) writeDescription(s string) {
	if s != "" {
		r.writeLine("")
		r.writeLine(strings.TrimSpace(s))
	}
}

func (r renderer) writeLine(s string) {
	_, err := r.WriteString(s + "\n")

	if err != nil {
		panic(err)
	}
}
