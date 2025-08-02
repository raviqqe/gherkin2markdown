package main_test

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/raviqqe/gherkin2markdown"
	"github.com/stretchr/testify/assert"
)

func TestCommandConvertDirectory(t *testing.T) {
	r, err := os.MkdirTemp("", "")
	assert.Nil(t, err)

	s := filepath.Join(r, "src")
	err = os.Mkdir(s, 0700)
	assert.Nil(t, err)

	err = os.WriteFile(filepath.Join(s, "foo.feature"), []byte("Feature: Foo"), 0600)
	assert.Nil(t, err)

	d := filepath.Join(r, "dest")

	assert.Nil(t, main.Run([]string{s, d}, io.Discard))

	bs, err := os.ReadFile(filepath.Join(d, "foo.md"))
	assert.Nil(t, err)
	assert.Equal(t, "# Foo\n", string(bs))

	assert.Nil(t, os.RemoveAll(r))
}
