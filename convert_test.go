package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertFile(t *testing.T) {
	f, err := ioutil.TempFile("", "")
	assert.Nil(t, err)
	defer os.Remove(f.Name())

	_, err = f.Write([]byte("Feature: Foo"))
	assert.Nil(t, err)

	assert.Nil(t, convertFile(f.Name(), ioutil.Discard))
}

func TestConvertFileError(t *testing.T) {
	f, err := ioutil.TempFile("", "")
	assert.Nil(t, err)
	defer os.Remove(f.Name())

	_, err = f.Write([]byte("Feature"))
	assert.Nil(t, err)

	assert.NotNil(t, convertFile(f.Name(), ioutil.Discard))
}

func TestConvertFilesWithNonReadableSourceDir(t *testing.T) {
	d, err := ioutil.TempDir("", "")
	assert.Nil(t, err)
	defer os.RemoveAll(d)

	assert.NotNil(t, convertFiles("foo", d))
}
