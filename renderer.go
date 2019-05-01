package main

import (
	"strings"

	"github.com/cucumber/cucumber-messages-go/v2"
	"github.com/willf/pad/utf8"
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

	if len(s.Examples) != 0 {
		r.writeLine("")
		r.renderExamples(s.Examples)
	}
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
		case *messages.Step_DataTable:
			r.renderDataTable(x.DataTable)
		default:
			panic("unreachable")
		}
	}
}

func (r renderer) renderExamples(es []*messages.Examples) {
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

func (r renderer) renderExampleTable(h *messages.TableRow, rs []*messages.TableRow) {
	ws := r.getCellWidths(append([]*messages.TableRow{h}, rs...))

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

func (r renderer) renderDataTable(t *messages.DataTable) {
	ws := r.getCellWidths(t.Rows)

	for _, t := range t.Rows {
		r.renderCells(t.Cells, ws)
	}
}

func (r renderer) renderCells(cs []*messages.TableCell, ws []int) {
	s := "|"

	for i, c := range cs {
		s += " " + utf8.Right(c.Value, ws[i], " ") + " |"
	}

	r.writeLine(s)
}

func (renderer) getCellWidths(rs []*messages.TableRow) []int {
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
