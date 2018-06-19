package main

import (
	"fmt"
	"strings"
	"testing"

	gherkin "github.com/cucumber/gherkin-go"
	"github.com/stretchr/testify/assert"
)

func TestNewRenderer(t *testing.T) {
	newRenderer()
}

func TestRendererRender(t *testing.T) {
	for _, ss := range [][2]string{
		{
			"Feature: Foo",
			"# Foo\n",
		},
		{`
Feature: Foo
  Scenario: Bar
    Given that
    When I do something
    Then something happens`, `
# Foo

## Bar

_Given_ that

_When_ I do something

_Then_ something happens`,
		},
		{`
Feature: Foo
  Scenario: Bar
    When I do something:
    """
    foo
    """`, fmt.Sprintf(`
# Foo

## Bar

_When_ I do something:

%[1]s
foo
%[1]s`, "```"),
		},
		{`
Feature: Foo

  bar`, `
# Foo

bar`,
		},
		{`
Feature: Foo
  Scenario: Bar

    baz`, `
# Foo

## Bar

baz`,
		},
	} {
		d, err := gherkin.ParseGherkinDocument(strings.NewReader(ss[0]))

		assert.Nil(t, err)
		assert.Equal(t, strings.TrimSpace(ss[1])+"\n", newRenderer().Render(d))
	}
}

func TestRendererRenderPanic(t *testing.T) {
	assert.Panics(t, func() {
		newRenderer().render(nil)
	})
}
