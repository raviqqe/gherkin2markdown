package main

import (
	"strings"

	"github.com/cucumber/gherkin-go"
	"github.com/willf/pad/utf8"
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
		case *gherkin.ScenarioOutline:
			r.renderScenarioOutline(x)
		default:
			panic("unreachable")
		}
	}
}

func (r renderer) renderBackground(b *gherkin.Background) {
	s := "## Background"

	if b.Name != "" {
		s += " (" + b.Name + ")"
	}

	r.writeLine(s)
	r.writeDescription(b.Description)
	r.renderSteps(b.Steps)
}

func (r renderer) renderScenario(s *gherkin.Scenario) {
	r.renderScenarioDefinition(&s.ScenarioDefinition)
}

func (r renderer) renderScenarioOutline(s *gherkin.ScenarioOutline) {
	r.renderScenarioDefinition(&s.ScenarioDefinition)

	if len(s.Examples) != 0 {
		r.writeLine("")
		r.renderExamples(s.Examples)
	}
}

func (r renderer) renderScenarioDefinition(s *gherkin.ScenarioDefinition) {
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
	r.writeLine("```" + d.ContentType)
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
		case *gherkin.DataTable:
			r.renderDataTable(x)
		default:
			panic("unreachable")
		}
	}
}

func (r renderer) renderExamples(es []*gherkin.Examples) {
	r.writeLine("### Examples")

	for _, e := range es {
		if e.Name != "" {
			r.writeLine("")
			r.writeLine("#### " + e.Name)
		}

		r.writeDescription(e.Description)

		r.writeLine("")
		r.renderExampleTable(e.TableHeader, e.TableBody)
	}
}

func (r renderer) renderExampleTable(h *gherkin.TableRow, rs []*gherkin.TableRow) {
	ws := r.getCellWidths(append([]*gherkin.TableRow{h}, rs...))

	r.renderCells(h.Cells, ws)

	s := "|"

	for _, w := range ws {
		s += strings.Repeat("-", w+2) + "|"
	}

	r.writeLine(s)

	for _, t := range rs {
		r.renderCells(t.Cells, ws)
	}
}

func (r renderer) renderDataTable(t *gherkin.DataTable) {
	ws := r.getCellWidths(t.Rows)

	for _, t := range t.Rows {
		r.renderCells(t.Cells, ws)
	}
}

func (r renderer) renderCells(cs []*gherkin.TableCell, ws []int) {
	s := "|"

	for i, c := range cs {
		s += " " + utf8.Right(c.Value, ws[i], " ") + " |"
	}

	r.writeLine(s)
}

func (renderer) getCellWidths(rs []*gherkin.TableRow) []int {
	ws := make([]int, len(rs[0].Cells))

	for _, r := range rs {
		for i, c := range r.Cells {
			if w := len(c.Value); w > ws[i] {
				ws[i] = w
			}
		}
	}

	return ws
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
