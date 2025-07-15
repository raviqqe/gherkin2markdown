package convert_test

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/raviqqe/gherkin2markdown/convert"
)

func TestFeatureFile(t *testing.T) {
	f, err := os.CreateTemp("", "")
	assert.Nil(t, err)
	defer func() { assert.Nil(t, os.Remove(f.Name())) }()

	_, err = f.Write([]byte("Feature: Foo"))
	assert.Nil(t, err)

	assert.Nil(t, convert.FeatureFile(f.Name(), io.Discard))
}

func TestFeatureFileError(t *testing.T) {
	f, err := os.CreateTemp("", "")
	assert.Nil(t, err)
	defer func() { assert.Nil(t, os.Remove(f.Name())) }()

	_, err = f.Write([]byte("Feature"))
	assert.Nil(t, err)

	assert.NotNil(t, convert.FeatureFile(f.Name(), io.Discard))
}

func TestFeatureFilesWithNonReadableSourceDir(t *testing.T) {
	d, err := os.MkdirTemp("", "")
	assert.Nil(t, err)
	defer func() { assert.Nil(t, os.RemoveAll(d)) }()

	assert.NotNil(t, convert.FeatureFiles("foo", d))
}
