package http_header_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/dbtedman/stop/internal/http_header"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func TestSetPermissionsPolicyHeader(t *testing.T) {
	response := &http.Response{Header: http.Header{}}
	http_header.SetPermissionsPolicyHeader(response)

	result := response.Header.Get(http_header.PermissionsPolicy)

	resultPartsPairs := strings.Split(result, ",")

	// reference:
	// - https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Permissions-Policy
	// - https://cheatsheetseries.owasp.org/cheatsheets/HTTP_Headers_Cheat_Sheet.html#permissions-policy-formerly-feature-policy
	allowedDirectives := []string{
		"camera",          // https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Permissions-Policy/camera
		"display-capture", // https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Permissions-Policy/display-capture
		"fullscreen",      // https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Permissions-Policy/fullscreen
		"geolocation",     // https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Permissions-Policy/geolocation
		"microphone",      // https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Permissions-Policy/microphone
		"web-share",       // https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Permissions-Policy/web-share
	}

	assert.GreaterOrEqual(t, len(resultPartsPairs), 1)

	for _, resultPartPair := range resultPartsPairs {
		resultParts := strings.Split(resultPartPair, "=")

		directive := strings.Trim(resultParts[0], " ")
		allowList := resultParts[1]

		assert.Truef(t, slices.Contains(allowedDirectives, directive), "[%s] is not an allowed directive, expected one of [%s]", directive, allowedDirectives)
		assert.Equal(t, "()", allowList, "allowList should be empty")
	}
}
