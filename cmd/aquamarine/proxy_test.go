package main

import (
	"errors"
	"github.com/dbtedman/stop/internal/http_header"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestNewProxyHandler(t *testing.T) {
	// given
	toUrl, _ := url.Parse("https://example.com")

	// when
	proxy := NewProxyHandler(*toUrl)

	// then
	assert.NotNil(t, proxy)
}

func TestNewProxyHandlerModifyResponse(t *testing.T) {
	// given
	toUrl, _ := url.Parse("https://example.com")
	proxy := NewProxyHandler(*toUrl)
	response := &http.Response{Header: http.Header{}}

	// when
	err := proxy.ModifyResponse(response)

	// then
	assert.Nil(t, err)
	assert.NotNil(t, response.Header.Get(PermissionsPolicy))
	assert.NotNil(t, response.Header.Get(ReferrerPolicy))
	assert.NotNil(t, response.Header.Get(XContentTypeOptions))
	assert.NotNil(t, response.Header.Get(XFrameOptions))
	assert.NotNil(t, response.Header.Get(http_header.StrictTransportSecurity))
	assert.NotNil(t, response.Header.Get(ContentSecurityPolicy))
}

func TestListenHTTPWithProxy(t *testing.T) {
	// given
	errorCh := make(chan error)
	proxy := TestProxy{}
	options, _ := NewOptions(":3001", "https://example.com")

	// when
	go func() {
		ListenHTTPWithProxy(&proxy, options, &errorCh)
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
	options, _ := NewOptions(":3001", "https://example.com")

	// when
	go func() {
		ListenHTTPWithProxy(&proxy, options, &errorCh)
	}()
	err := <-errorCh

	// then
	assert.ErrorContains(t, err, alwaysErrorMessage)
}

const alwaysErrorMessage = "always showing error"

type alwaysErrorProxy struct {
}

var _ Proxy = &alwaysErrorProxy{}

func (t *alwaysErrorProxy) ListenAndServe(addr string, handler http.HandlerFunc) error {
	return errors.New(alwaysErrorMessage)
}
