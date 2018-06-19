package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetArguments(t *testing.T) {
	for _, c := range []struct {
		parameters []string
		arguments
	}{
		{[]string{"file"}, arguments{File: "file"}},
		{[]string{"dir1", "dir2"}, arguments{SrcDir: "dir1", DestDir: "dir2"}},
	} {
		assert.Equal(t, c.arguments, getArguments(c.parameters))
	}
}

func TestParseArgumentsPanic(t *testing.T) {
	assert.Panics(t, func() {
		parseArguments("", []string{"file"}, &arguments{})
	})

	assert.Panics(t, func() {
		parseArguments(usage, []string{"file"}, arguments{})
	})
}
