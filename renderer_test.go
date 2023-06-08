package main

import (
	"fmt"
	"strings"
	"testing"

	gherkin "github.com/cucumber/gherkin-go/v19"
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

_Then_ something happens.`,
		},
		{`
Feature: Foo
  Scenario: Bar
    When I do something:
    """sh
    foo
    """`, fmt.Sprintf(`
# Foo

## Bar

_When_ I do something:

%[1]ssh
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
		{`
Feature: Foo
  Background: Bar
    When I do something`, `
# Foo

## Background (Bar)

_When_ I do something.`,
		},
		{`
Feature: Foo
  Background: Bar
  Given Baz:
    | foo |
    | bar |`, `
# Foo

## Background (Bar)

_Given_ Baz:

| foo |
| bar |`,
		},
		{`
Feature: Foo
  Scenario Outline: Bar
    When <someone> does <something>.
    Examples:
      | someone | something |
      | I       | cooking   |
      | You     | coding    |`, `
# Foo

## Bar

_When_ <someone> does <something>.

### Examples

| someone | something |
|---------|-----------|
| I       | cooking   |
| You     | coding    |`},
		{`
Feature: Foo
  Scenario Outline: Bar
    When <someone> does <something>.
    Examples: Baz
      | someone | something |
      | I       | cooking   |
      | You     | coding    |`, `
# Foo

## Bar

_When_ <someone> does <something>.

### Examples

#### Baz

| someone | something |
|---------|-----------|
| I       | cooking   |
| You     | coding    |`},
		{`
Feature: Foo
  Scenario Outline: Bar
    When <someone> does <something>.
    Examples: Baz
      foo bar baz.

      | someone | something |
      | I       | cooking   |
      | You     | coding    |`, `
# Foo

## Bar

_When_ <someone> does <something>.

### Examples

#### Baz

foo bar baz.

| someone | something |
|---------|-----------|
| I       | cooking   |
| You     | coding    |`},
		{`
Feature: Foo
  Scenario Outline: Bar
    When <someone> does <something>.
    Examples: Baz
      | someone |
      | I       |
      | You     |
    Examples: Blah
      | something |
      | cooking   |
      | coding    |`, `
# Foo

## Bar

_When_ <someone> does <something>.

### Examples

#### Baz

| someone |
|---------|
| I       |
| You     |

#### Blah

| something |
|-----------|
| cooking   |
| coding    |`}, {`
Feature: Foo
	Rule: Bar
		Example: Baz
			When qux
`, `
# Foo

## Bar

### Baz

_When_ qux.
`},
	} {
		d, err := gherkin.ParseGherkinDocument(strings.NewReader(ss[0]), func() string { return "" })

		assert.Nil(t, err)
		assert.Equal(t, strings.TrimSpace(ss[1])+"\n", newRenderer().Render(d))
	}
}
