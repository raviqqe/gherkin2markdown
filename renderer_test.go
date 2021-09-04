package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"

	gherkin "github.com/cucumber/gherkin-go/v19"
)

func TestNewRenderer(t *testing.T) {
	newRenderer()
}

func TestRendererRender(t *testing.T) {

	testTable:= []struct {
		Input string
		Expected string
	}{
		{
			"Feature: Foo",
			`
---
layout: default
title: Foo
parent: Features
---
# Foo`,
		},
		{`
Feature: Foo
  Scenario: Bar
    Given that
    When I do something
    Then something happens`, `
---
layout: default
title: Foo
parent: Features
---
# Foo

## Scenarios

### Bar

**Given** that
**When** I do something
**Then** something happens.`,
		},
		{`
Feature: Foo
  Scenario: Bar
    When I do something:
    """sh
    foo
    """`, fmt.Sprintf(`
---
layout: default
title: Foo
parent: Features
---
# Foo

## Scenarios

### Bar

**When** I do something:

%[1]ssh
foo
%[1]s`, "```"),
		},
		{`
Feature: Foo

  bar`, `
---
layout: default
title: Foo
parent: Features
---
# Foo

bar`,
		},
		{`
Feature: Foo
  Scenario: Bar

    baz`, `
---
layout: default
title: Foo
parent: Features
---
# Foo

## Scenarios

### Bar


baz`,
		},
		{`
Feature: Foo
  Background: Bar
    When I do something`, `
---
layout: default
title: Foo
parent: Features
---
# Foo

## Background (Bar)

**When** I do something.`,
		},
		{`
Feature: Foo
  Background: Bar
  Given Baz:
    | foo |
    | bar |`, `
---
layout: default
title: Foo
parent: Features
---
# Foo

## Background (Bar)

**Given** Baz:

| foo |
|:----|
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
---
layout: default
title: Foo
parent: Features
---
# Foo

## Scenarios

### Bar

**When** <someone> does <something>.

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
---
layout: default
title: Foo
parent: Features
---
# Foo

## Scenarios

### Bar

**When** <someone> does <something>.

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
---
layout: default
title: Foo
parent: Features
---
# Foo

## Scenarios

### Bar

**When** <someone> does <something>.

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
---
layout: default
title: Foo
parent: Features
---
# Foo

## Scenarios

### Bar

**When** <someone> does <something>.

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
| coding    |`},
	}

	for i, ss := range testTable {
		t.Run(fmt.Sprintf("Render test %d",i), func(t *testing.T) {
			d, err := gherkin.ParseGherkinDocument(strings.NewReader(ss.Input), func() string { return "" })

			assert.Nil(t, err)
			assert.Equal(t, strings.TrimSpace(ss.Expected), strings.TrimSpace(newRenderer().Render(d)))
		})
	}
}
