# gherkin2markdown

[![Circle CI](https://img.shields.io/circleci/project/github/raviqqe/gherkin2markdown/master.svg?style=flat-square)](https://circleci.com/gh/raviqqe/gherkin2markdown)
[![Codecov](https://img.shields.io/codecov/c/github/raviqqe/gherkin2markdown.svg?style=flat-square)](https://codecov.io/gh/raviqqe/gherkin2markdown)
[![Go Report Card](https://goreportcard.com/badge/github.com/raviqqe/gherkin2markdown?style=flat-square)](https://goreportcard.com/report/github.com/raviqqe/gherkin2markdown)
[![License](https://img.shields.io/github/license/raviqqe/gherkin2markdown.svg?style=flat-square)](LICENSE)

A command to convert Gherkin files into Markdown.

## Installation

```
go get -u github.com/raviqqe/gherkin2markdown
```

## Usage

```
gherkin2markdown <file>
```

or

```
gherkin2markdown <srcdir> <destdir>
```

## Example

Given a file named `math.feature` with:

```gherkin
Feature: Math
  Scenario: Add 2 numbers
    Given a file named "main.cloe" with:
    """
    (write (+ 2016 33))
    """
    When I successfully run `cloe main.cloe`
    Then the stdout should contain exactly "2049"

  Scenario: Subtract a number from the other
    Given a file named "main.cloe" with:
    """
    (write (- 2049 33))
    """
    When I successfully run `cloe main.cloe`
    Then the stdout should contain exactly "2016"
```

When I successfully run `gherkin2markdown math.feature`

Then the stdout should contain exactly:

````markdown
# Math

## Add 2 numbers

_Given_ a file named "main.cloe" with:

```
(write (+ 2016 33))
```

_When_ I successfully run `cloe main.cloe`

_Then_ the stdout should contain exactly "2049"

## Subtract a number from the other

_Given_ a file named "main.cloe" with:

```
(write (- 2049 33))
```

_When_ I successfully run `cloe main.cloe`

_Then_ the stdout should contain exactly "2016"
````

## License

[MIT](LICENSE)
