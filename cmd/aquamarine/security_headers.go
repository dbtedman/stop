package main

import "net/http"

const ReferrerPolicy = "Referrer-Policy"
const XContentTypeOptions = "X-Content-Type-Options"
const ContentSecurityPolicy = "Content-Security-Policy"

func SetContentSecurityPolicyHeader(response *http.Response) {
	// Start by denying all content, then selectively allowing it.
	defaultSrc := "default-src 'none'"

	// Allow scripts from the same location as the html.
	scriptSrc := "script-src 'self'"

	// Allow ajax requests to the same location as the html.
	connectSrc := "connect-src 'self'"

	// Allow images from the same location as the html.
	imageSrc := "img-src 'self'"

	// Allow style sheets from the same location as the html.
	styleSrc := "style-src 'self'"

	fontSrc := "font-src 'self'"

	// Prevent all framing of the content.
	frameAncestors := "frame-ancestors 'none'"

	// Allow form submissions only to the same location as the html.
	formAction := "form-action 'self'"

	policy := defaultSrc + "; " + scriptSrc + "; " + connectSrc + "; " + imageSrc + "; " + styleSrc + "; " + fontSrc + "; " + frameAncestors + "; " + formAction + ";"

	response.Header.Set("Content-Security-Policy", policy)
}

func SetXContentTypeOptionsHeader(response *http.Response) {
	response.Header.Set(XContentTypeOptions, "nosniff")
}

func SetReferrerPolicyHeader(response *http.Response) {
	response.Header.Set(ReferrerPolicy, "strict-origin-when-cross-origin")
}
