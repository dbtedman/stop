package http_header_test

import (
	"github.com/dbtedman/stop/internal/http_header"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestSetStrictTransportSecurityHeader(t *testing.T) {
	response := &http.Response{Header: http.Header{}}
	http_header.SetStrictTransportSecurityHeader(response)

	result := response.Header.Get(http_header.StrictTransportSecurity)

	// reference:
	// - https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security
	// - https://cheatsheetseries.owasp.org/cheatsheets/HTTP_Headers_Cheat_Sheet.html#strict-transport-security-hsts
	// - https://hstspreload.org
	assert.Equal(t, result, "max-age=300; includeSubDomains")
}
