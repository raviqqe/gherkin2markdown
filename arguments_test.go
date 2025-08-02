package main_test

import (
	"testing"

	"github.com/raviqqe/gherkin2markdown"
	"github.com/stretchr/testify/assert"
)

func TestGetArguments(t *testing.T) {
	for _, c := range []struct {
		parameters []string
		arguments  main.Arguments
	}{
		{[]string{"dir1", "dir2"}, main.Arguments{SrcDir: "dir1", DestDir: "dir2"}},
	} {
		args, err := main.GetArguments(c.parameters)

		assert.Nil(t, err)
		assert.Equal(t, c.arguments, args)
	}
}
