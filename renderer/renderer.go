package renderer

import (
	"strings"

	messages "github.com/cucumber/messages/go/v24"
	"github.com/willf/pad/utf8"
)

// Renderer represents various parts of a Cucumber document as Markdown
// strings.
type Renderer interface {
	// Render returns a Markdown string for the given
	// [*messages.GherkinDocument].
	Render(d *messages.GherkinDocument) string
}

type renderer struct {
	*strings.Builder
	depth int
}

// New returns a new [Renderer].
func New() Renderer {
	return &renderer{&strings.Builder{}, 0}
}

// Render returns a Markdown string for the given [*messages.GherkinDocument].
func (r *renderer) Render(d *messages.GherkinDocument) string {
	r.renderFeature(d.Feature)

	return r.String()
}

func (r *renderer) renderFeature(f *messages.Feature) {
	r.writeHeadline(f.Name)

	r.depth++
	defer func() { r.depth-- }()

	r.writeDescription(f.Description)

	for _, c := range f.Children {
		r.writeLine("")

		if c.Background != nil {
			r.renderBackground(c.Background)
		}

		if c.Scenario != nil {
			r.renderScenario(c.Scenario)
		}

		if c.Rule != nil {
			r.renderRule(c.Rule)
		}
	}
}

func (r *renderer) renderBackground(b *messages.Background) {
	s := "Background"

	if b.Name != "" {
		s += " (" + b.Name + ")"
	}

	r.writeHeadline(s)

	r.depth++
	defer func() { r.depth-- }()

	r.writeDescription(b.Description)
	r.renderSteps(b.Steps)
}

func (r *renderer) renderScenario(s *messages.Scenario) {
	r.writeHeadline(s.Name)

	r.depth++
	defer func() { r.depth-- }()

	r.writeDescription(s.Description)
	r.renderSteps(s.Steps)

	if len(s.Examples) != 0 {
		r.writeLine("")
		r.renderExamples(s.Examples)
	}
}

func (r *renderer) renderRule(l *messages.Rule) {
	r.writeHeadline(l.Name)

	r.depth++
	defer func() { r.depth-- }()

	r.writeDescription(l.Description)

	for _, c := range l.Children {
		r.writeLine("")

		if c.Background != nil {
			r.renderBackground(c.Background)
		}

		if c.Scenario != nil {
			r.renderScenario(c.Scenario)
		}
	}
}

func (r *renderer) renderSteps(ss []*messages.Step) {
	for i, s := range ss {
		r.writeLine("")
		r.renderStep(s, i == len(ss)-1)
	}
}

func (r *renderer) renderDocString(d *messages.DocString) {
	r.writeLine("```" + d.MediaType)
	r.writeLine(d.Content)
	r.writeLine("```")
}

func (r *renderer) renderStep(s *messages.Step, last bool) {
	if last && s.DocString == nil && s.DataTable == nil && s.Text[len(s.Text)-1] != '.' {
		s.Text += "."
	}

	text := strings.ReplaceAll(s.Text, "<", `\<`)
	text = strings.ReplaceAll(text, ">", `\>`)

	r.writeLine("_" + strings.TrimSpace(s.Keyword) + "_ " + text)

	if s.DocString != nil {
		r.writeLine("")
		r.renderDocString(s.DocString)
	}

	if s.DataTable != nil {
		r.writeLine("")
		r.renderDataTable(s.DataTable)
	}
}

func (r *renderer) renderExamples(es []*messages.Examples) {
	r.writeHeadline("Examples")

	r.depth++
	defer func() { r.depth-- }()

	for _, e := range es {
		if e.Name != "" {
			r.writeLine("")
			r.writeHeadline(e.Name)
		}

		r.writeDescription(e.Description)

		r.writeLine("")
		r.renderTable(e.TableHeader, e.TableBody)
	}
}

func (r renderer) renderTable(h *messages.TableRow, rs []*messages.TableRow) {
	ws := r.getCellWidths(append([]*messages.TableRow{h}, rs...))

	if h == nil {
		empty := make([]*messages.TableCell, len(ws))
		r.renderCells(empty, ws)
	} else {
		r.renderCells(h.Cells, ws)
	}

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
	r.renderTable(nil, t.Rows)
}

func (r renderer) renderCells(cs []*messages.TableCell, ws []int) {
	s := "|"
	cn := len(cs)

	for i := range ws {
		v := ""

		if cn > i && cs[i] != nil {
			v = cs[i].Value
		}

		s += " " + utf8.Right(v, ws[i], " ") + " |"
	}

	r.writeLine(s)
}

func (renderer) getCellWidths(rs []*messages.TableRow) []int {
	cols := 0

	for _, r := range rs {
		if r != nil {
			if n := len(r.Cells); n > cols {
				cols = n
			}
		}
	}

	ws := make([]int, cols)

	for _, r := range rs {
		if r != nil {
			for i, c := range r.Cells {
				if w := len(c.Value); w > ws[i] {
					ws[i] = w
				}
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

func (r renderer) writeHeadline(s string) {
	r.writeLine(strings.Repeat("#", r.depth+1) + " " + s)
}

func (r renderer) writeLine(s string) {
	_, err := r.WriteString(s + "\n")

	if err != nil {
		panic(err)
	}
}
