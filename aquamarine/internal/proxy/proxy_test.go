package proxy_test

import (
	"net/http"
	"net/url"
	"testing"

	application2 "github.com/dbtedman/stop/aquamarine/internal/proxy"
	"github.com/dbtedman/stop/aquamarine/internal/security"

	"github.com/stretchr/testify/assert"
)

func TestNewProxyHandler(t *testing.T) {
	// given
	toUrl, _ := url.Parse("https://example.com")

	// when
	proxy := application2.NewProxyHandler(*toUrl)

	// then
	assert.NotNil(t, proxy)
}

func TestNewProxyHandlerModifyResponse(t *testing.T) {
	// given
	toUrl, _ := url.Parse("https://example.com")
	proxy := application2.NewProxyHandler(*toUrl)
	response := &http.Response{Header: http.Header{}}

	// when
	err := proxy.ModifyResponse(response)

	// then
	assert.Nil(t, err)
	assert.NotNil(t, response.Header.Get(security.PermissionsPolicy))
	assert.NotNil(t, response.Header.Get(security.ReferrerPolicy))
	assert.NotNil(t, response.Header.Get(security.XContentTypeOptions))
	assert.NotNil(t, response.Header.Get(security.XFrameOptions))
	assert.NotNil(t, response.Header.Get(security.StrictTransportSecurity))
	assert.NotNil(t, response.Header.Get(security.ContentSecurityPolicy))
}
