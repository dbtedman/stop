package cmd_test

import (
	"bytes"
	"testing"

	"github.com/dbtedman/stop/aquamarine/cmd"
	"github.com/dbtedman/stop/aquamarine/internal/proxy"

	"github.com/stretchr/testify/assert"
)

func TestServeCommand(t *testing.T) {
	// given
	errorCh := make(chan error)
	var errConsole bytes.Buffer
	var outConsole bytes.Buffer
	proxy := proxy.TestProxy{}
	command := cmd.ServeCommand(&proxy, &errorCh)
	command.SetErr(&errConsole)
	command.SetOut(&outConsole)

	// when
	go func() {
		command.SetArgs([]string{"", "--from=:3000", "--to=https://example.com"})
		errorCh <- command.Execute()
	}()
	err := <-errorCh

	// then
	assert.Nil(t, err)
	assert.Equal(t, "", errConsole.String())
	assert.Equal(t, ":3000", proxy.Addr)
}
