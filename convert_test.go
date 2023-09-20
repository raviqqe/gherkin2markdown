package main_test

import (
	"io"
	"os"
	"testing"

	"github.com/raviqqe/gherkin2markdown"
	"github.com/stretchr/testify/assert"
)

func TestConvertFile(t *testing.T) {
	f, err := os.CreateTemp("", "")
	assert.Nil(t, err)
	defer os.Remove(f.Name())

	_, err = f.Write([]byte("Feature: Foo"))
	assert.Nil(t, err)

	assert.Nil(t, main.ConvertFile(f.Name(), io.Discard))
}

func TestConvertFileError(t *testing.T) {
	f, err := os.CreateTemp("", "")
	assert.Nil(t, err)
	defer os.Remove(f.Name())

	_, err = f.Write([]byte("Feature"))
	assert.Nil(t, err)

	assert.NotNil(t, main.ConvertFile(f.Name(), io.Discard))
}

func TestConvertFilesWithNonReadableSourceDir(t *testing.T) {
	d, err := os.MkdirTemp("", "")
	assert.Nil(t, err)
	defer os.RemoveAll(d)

	assert.NotNil(t, main.ConvertFiles("foo", d))
}
