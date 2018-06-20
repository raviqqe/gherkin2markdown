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
	r.renderFeature(d.Feature)

	return r.Builder.String()
}

func (r renderer) renderFeature(f *gherkin.Feature) {
	r.writeLine("# " + f.Name)
	r.writeDescription(f.Description)

	for _, x := range f.Children {
		r.writeLine("")

		switch x := x.(type) {
		case *gherkin.Background:
			r.renderBackground(x)
		case *gherkin.Scenario:
			r.renderScenario(x)
		default:
			panic("unreachable")
		}
	}
}

func (r renderer) renderBackground(b *gherkin.Background) {
	r.writeLine("## Background (" + b.Name + ")")
	r.writeDescription(b.Description)
	r.renderSteps(b.Steps)
}

func (r renderer) renderScenario(s *gherkin.Scenario) {
	r.writeLine("## " + s.Name)
	r.writeDescription(s.Description)
	r.renderSteps(s.Steps)
}

func (r renderer) renderSteps(ss []*gherkin.Step) {
	for i, s := range ss {
		r.writeLine("")
		r.renderStep(s, i == len(ss)-1)
	}
}

func (r renderer) renderDocString(d *gherkin.DocString) {
	r.writeLine("```")
	r.writeLine(d.Content)
	r.writeLine("```")
}

func (r renderer) renderStep(s *gherkin.Step, last bool) {
	if last && s.Argument == nil && s.Text[len(s.Text)-1] != '.' {
		s.Text += "."
	}

	r.writeLine("_" + strings.TrimSpace(s.Keyword) + "_ " + s.Text)

	if s.Argument != nil {
		r.writeLine("")

		switch x := s.Argument.(type) {
		case *gherkin.DocString:
			r.renderDocString(x)
		default:
			panic("unreachable")
		}
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
