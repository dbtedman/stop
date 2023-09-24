package http_header_test

import (
	"net/http"
	"testing"

	"github.com/dbtedman/stop/internal/http_header"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func TestSetXFrameOptionsHeader(t *testing.T) {
	response := &http.Response{Header: http.Header{}}
	http_header.SetXFrameOptionsHeader(response)

	result := response.Header.Get(http_header.XFrameOptions)

	// reference:
	// - https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options
	// - https://cheatsheetseries.owasp.org/cheatsheets/HTTP_Headers_Cheat_Sheet.html#x-frame-options
	allowedResults := []string{
		"DENY",
		"SAMEORIGIN",
	}
	assert.Truef(t, slices.Contains(allowedResults, result), "result [%s] is not an allowed value [%s]", result, allowedResults)
}
