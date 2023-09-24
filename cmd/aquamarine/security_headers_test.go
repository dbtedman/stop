package main

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
	"net/http"
	"strings"
	"testing"
)

func TestSetContentSecurityPolicyHeader(t *testing.T) {
	response := &http.Response{Header: http.Header{}}
	SetContentSecurityPolicyHeader(response)

	result := response.Header.Get(ContentSecurityPolicy)

	// reference:
	// - https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy
	// - https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP
	// - https://cheatsheetseries.owasp.org/cheatsheets/HTTP_Headers_Cheat_Sheet.html#content-security-policy-csp
	// - https://cheatsheetseries.owasp.org/cheatsheets/Content_Security_Policy_Cheat_Sheet.html
	// - https://csp-evaluator.withgoogle.com

	resultPartsPairs := strings.Split(result, ";")

	allowedDirectives := []string{
		"default-src",
		"connect-src",     // https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/connect-src
		"font-src",        // https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/font-src
		"form-action",     // https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/form-action
		"frame-ancestors", // https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/frame-ancestors
		"img-src",         // https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/img-src
		"script-src",      // https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/script-src
		"style-src",       // https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/style-src
	}

	assert.GreaterOrEqual(t, len(resultPartsPairs), 1)

	for _, resultPartPair := range resultPartsPairs {
		resultParts := strings.Split(strings.Trim(resultPartPair, " "), " ")
		directive := strings.Trim(resultParts[0], " ")

		if directive == "" {
			// Ignore last semicolon
			continue
		}

		assert.Truef(t, slices.Contains(allowedDirectives, directive), "[%s] is not an allowed directive, expected one of [%s]", directive, allowedDirectives)
	}
}

func TestSetXFrameOptionsHeader(t *testing.T) {
	response := &http.Response{Header: http.Header{}}
	SetXFrameOptionsHeader(response)

	result := response.Header.Get(XFrameOptions)

	// reference:
	// - https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options
	// - https://cheatsheetseries.owasp.org/cheatsheets/HTTP_Headers_Cheat_Sheet.html#x-frame-options
	allowedResults := []string{
		"DENY",
		"SAMEORIGIN",
	}
	assert.Truef(t, slices.Contains(allowedResults, result), "result [%s] is not an allowed value [%s]", result, allowedResults)
}

func TestSetXContentTypeOptionsHeader(t *testing.T) {
	response := &http.Response{Header: http.Header{}}
	SetXContentTypeOptionsHeader(response)

	result := response.Header.Get(XContentTypeOptions)

	// reference:
	// - https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options
	// - https://cheatsheetseries.owasp.org/cheatsheets/HTTP_Headers_Cheat_Sheet.html
	assert.Equal(t, result, "nosniff")
}

func TestSetReferrerPolicyHeader(t *testing.T) {
	response := &http.Response{Header: http.Header{}}
	SetReferrerPolicyHeader(response)

	result := response.Header.Get(ReferrerPolicy)

	// reference:
	// - https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy
	// - https://cheatsheetseries.owasp.org/cheatsheets/HTTP_Headers_Cheat_Sheet.html#referrer-policy
	allowedResults := []string{
		"strict-origin-when-cross-origin",
		"no-referrer",
		"no-referrer-when-downgrade",
		"origin",
		"origin-when-cross-origin",
		"same-origin",
		"strict-origin",
		"strict-origin-when-cross-origin",
		"unsafe-url",
	}
	assert.Truef(t, slices.Contains(allowedResults, result), "result [%s] is not an allowed value [%s]", result, allowedResults)
}

func TestSetPermissionsPolicyHeader(t *testing.T) {
	response := &http.Response{Header: http.Header{}}
	SetPermissionsPolicyHeader(response)

	result := response.Header.Get(PermissionsPolicy)

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
