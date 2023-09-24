package http_header

import "net/http"

const PermissionsPolicy = "Permissions-Policy"

func SetPermissionsPolicyHeader(response *http.Response) {
	response.Header.Set(PermissionsPolicy, "geolocation=(), camera=(), microphone=(), display-capture=()")
}
