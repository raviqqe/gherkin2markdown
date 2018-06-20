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
		r.render(x.Steps)
	case *gherkin.Scenario:
		r.writeLine("## " + x.Name)
		r.writeDescription(x.Description)
		r.render(x.Steps)
	case []*gherkin.Step:
		for i, s := range x {
			r.writeLine("")
			r.renderStep(s, i == len(x)-1)
		}
	case *gherkin.DocString:
		r.writeLine("```")
		r.writeLine(x.Content)
		r.writeLine("```")
	default:
		panic("unreachable")
	}
}

func (r renderer) renderStep(s *gherkin.Step, last bool) {
	if last && s.Argument == nil && s.Text[len(s.Text)-1] != '.' {
		s.Text += "."
	}

	r.writeLine("_" + strings.TrimSpace(s.Keyword) + "_ " + s.Text)

	if s.Argument != nil {
		r.writeLine("")
		r.render(s.Argument)
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
