package main

import (
	"strings"

	messages "github.com/cucumber/cucumber-messages-go"
)

type renderer struct {
	*strings.Builder
}

func newRenderer() renderer {
	return renderer{&strings.Builder{}}
}

func (r renderer) Render(d *messages.GherkinDocument) string {
	r.renderFeature(d.Feature)

	return r.Builder.String()
}

func (r renderer) renderFeature(f *messages.Feature) {
	r.writeLine("# " + f.Name)
	r.writeDescription(f.Description)

	for _, x := range f.Children {
		r.writeLine("")

		switch x := x.Value.(type) {
		case *messages.FeatureChild_Background:
			r.renderBackground(x.Background)
		case *messages.FeatureChild_Scenario:
			r.renderScenario(x.Scenario)
		default:
			panic("unreachable")
		}
	}
}

func (r renderer) renderBackground(b *messages.Background) {
	r.writeLine("## Background (" + b.Name + ")")
	r.writeDescription(b.Description)
	r.renderSteps(b.Steps)
}

func (r renderer) renderScenario(s *messages.Scenario) {
	r.writeLine("## " + s.Name)
	r.writeDescription(s.Description)
	r.renderSteps(s.Steps)
}

func (r renderer) renderSteps(ss []*messages.Step) {
	for i, s := range ss {
		r.writeLine("")
		r.renderStep(s, i == len(ss)-1)
	}
}

func (r renderer) renderDocString(d *messages.DocString) {
	r.writeLine("```")
	r.writeLine(d.Content)
	r.writeLine("```")
}

func (r renderer) renderStep(s *messages.Step, last bool) {
	if last && s.Argument == nil && s.Text[len(s.Text)-1] != '.' {
		s.Text += "."
	}

	r.writeLine("_" + strings.TrimSpace(s.Keyword) + "_ " + s.Text)

	if s.Argument != nil {
		r.writeLine("")

		switch x := s.Argument.(type) {
		case *messages.Step_DocString:
			r.renderDocString(x.DocString)
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
