package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionCommand(t *testing.T) {
	// given
	errorCh := make(chan error)
	var errConsole bytes.Buffer
	var outConsole bytes.Buffer
	command := VersionCommand(&errorCh)
	command.SetErr(&errConsole)
	command.SetOut(&outConsole)

	// when
	go func() {
		errorCh <- command.Execute()
	}()
	err := <-errorCh

	// then
	assert.Nil(t, err)
	assert.Equal(t, "", errConsole.String())
	assert.Contains(t, outConsole.String(), "Conveyance version: latest")
	assert.Contains(t, outConsole.String(), "commit: n/a")
	assert.Contains(t, outConsole.String(), "built at: n/a")
}
