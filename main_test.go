package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {
	f, err := ioutil.TempFile("", "")
	assert.Nil(t, err)

	f.WriteString("Feature: Foo")

	assert.Nil(t, command([]string{f.Name()}, ioutil.Discard))

	os.Remove(f.Name())
}

func TestCommandWithNonExistentFile(t *testing.T) {
	assert.NotNil(t, command([]string{"non-existent.feature"}, ioutil.Discard))
}

func TestCommandWithDirectory(t *testing.T) {
	r, err := ioutil.TempDir("", "")
	assert.Nil(t, err)

	s := filepath.Join(r, "src")
	err = os.Mkdir(s, 0700)
	assert.Nil(t, err)

	err = ioutil.WriteFile(filepath.Join(s, "foo.feature"), []byte("Feature: Foo"), 0600)
	assert.Nil(t, err)

	d := filepath.Join(r, "dest")

	assert.Nil(t, command([]string{s, d}, ioutil.Discard))

	bs, err := ioutil.ReadFile(filepath.Join(d, "foo.md"))
	assert.Nil(t, err)
	assert.Equal(t, "# Foo\n", string(bs))

	os.RemoveAll(r)
}
