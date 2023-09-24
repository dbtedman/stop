package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const aListenAddress = "https://localhost"
const aProxyAddress = "http://localhost:8080"

func TestNewOptions(t *testing.T) {
	theOptions, err := NewOptions(
		aListenAddress,
		aProxyAddress,
	)

	assert.Nil(t, err)
	assert.NotEmpty(t, theOptions.ListenAddress())
	assert.NotEmpty(t, theOptions.ProxyAddressURL())
}

func TestNewOptionsRejectsAbsentListenAddress(t *testing.T) {
	_, err := NewOptions(
		"",
		aProxyAddress,
	)

	assert.Error(t, err)
}

func TestNewOptionsRejectsAbsentProxyAddress(t *testing.T) {
	_, err := NewOptions(
		aListenAddress,
		"",
	)

	assert.Error(t, err)
}

func TestNewOptionsRejectsInvalidProxyAddress(t *testing.T) {
	_, err := NewOptions(
		aListenAddress,
		"://localhost",
	)

	assert.Error(t, err)
}
