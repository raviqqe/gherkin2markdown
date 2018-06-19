package main

import (
	"io/ioutil"
	"os"
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
