package web_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/dbtedman/stop/aquamarine/internal/options"
	"github.com/dbtedman/stop/aquamarine/internal/proxy"
	"github.com/dbtedman/stop/aquamarine/web"

	"github.com/stretchr/testify/assert"
)

func TestListenHTTPWithProxy(t *testing.T) {
	// given
	errorCh := make(chan error)
	proxy := proxy.TestProxy{}
	options, _ := options.NewOptions(":3001", "https://example.com")

	// when
	go func() {
		web.ListenHTTPWithProxy(&proxy, options, &errorCh)
	}()
	err := <-errorCh

	// then
	assert.Nil(t, err)
	assert.Equal(t, ":3001", proxy.Addr)
}

func TestListenHTTPWithProxyHandlesListenAndServeError(t *testing.T) {
	// given
	errorCh := make(chan error)
	proxy := alwaysErrorProxy{}
	options, _ := options.NewOptions(":3001", "https://example.com")

	// when
	go func() {
		web.ListenHTTPWithProxy(&proxy, options, &errorCh)
	}()
	err := <-errorCh

	// then
	assert.ErrorContains(t, err, alwaysErrorMessage)
}

const alwaysErrorMessage = "always showing error"

type alwaysErrorProxy struct {
}

var _ proxy.Proxy = &alwaysErrorProxy{}

func (t *alwaysErrorProxy) ListenAndServe(addr string, handler http.HandlerFunc) error {
	return errors.New(alwaysErrorMessage)
}
