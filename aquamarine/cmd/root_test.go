package cmd_test

import (
	"bytes"
	"testing"

	"github.com/dbtedman/stop/aquamarine/cmd"

	"github.com/stretchr/testify/assert"
)

func TestRootCommand(t *testing.T) {
	// given
	errorCh := make(chan error)
	var errConsole bytes.Buffer
	var outConsole bytes.Buffer
	command := cmd.RootCommand(&errorCh)
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
	assert.Contains(t, outConsole.String(), "Provide security by proxying requests to legacy applications.")
	assert.Contains(t, outConsole.String(), "-h, --help   help for conveyance")
}
