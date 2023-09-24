package security

import "net/http"

const PermissionsPolicy = "Permissions-Policy"
const ReferrerPolicy = "Referrer-Policy"
const XContentTypeOptions = "X-Content-Type-Options"
const XFrameOptions = "X-Frame-Options"
const StrictTransportSecurity = "Strict-Transport-Security"
const ContentSecurityPolicy = "Content-Security-Policy"

func SetStrictTransportSecurityHeader(response *http.Response) {
	response.Header.Set(StrictTransportSecurity, "max-age=300; includeSubDomains")
}

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

func SetXFrameOptionsHeader(response *http.Response) {
	response.Header.Set(XFrameOptions, "DENY")
}

func SetXContentTypeOptionsHeader(response *http.Response) {
	response.Header.Set(XContentTypeOptions, "nosniff")
}

func SetReferrerPolicyHeader(response *http.Response) {
	response.Header.Set(ReferrerPolicy, "strict-origin-when-cross-origin")
}

func SetPermissionsPolicyHeader(response *http.Response) {
	response.Header.Set(PermissionsPolicy, "geolocation=(), camera=(), microphone=(), display-capture=()")
}
