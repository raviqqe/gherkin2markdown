# gherkin2markdown

[![GitHub Action](https://img.shields.io/github/actions/workflow/status/raviqqe/gherkin2markdown/main.yaml?branch=main&style=flat-square)](https://github.com/raviqqe/gherkin2markdown/actions)
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

_Given_ a file named `math.feature` with:

```gherkin
Feature: Python
  Scenario: Hello, world!
    Given a file named "main.py" with:
    """
    print("Hello, world!")
    """
    When I successfully run `python3 main.py`
    Then the stdout should contain exactly "Hello, world!"

  Scenario Outline: Add numbers
    Given a file named "main.py" with:
    """
    print(<x> + <y>)
    """
    When I successfully run `python3 main.py`
    Then the stdout should contain exactly "<z>"
    Examples:
      | x | y | z |
      | 1 | 2 | 3 |
      | 4 | 5 | 9 |
```

_When_ I successfully run `gherkin2markdown math.feature`

_Then_ the stdout should contain exactly:

````markdown
# Python

## Hello, world!

_Given_ a file named "main.py" with:

```
print("Hello, world!")
```

_When_ I successfully run `python3 main.py`

_Then_ the stdout should contain exactly "Hello, world!".

## Add numbers

_Given_ a file named "main.py" with:

```
print(<x> + <y>)
```

_When_ I successfully run `python3 main.py`

_Then_ the stdout should contain exactly "<z>".

### Examples

| x   | y   | z   |
| --- | --- | --- |
| 1   | 2   | 3   |
| 4   | 5   | 9   |
````

## License

[MIT](LICENSE)
