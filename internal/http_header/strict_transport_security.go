package http_header

import "net/http"

const StrictTransportSecurity = "Strict-Transport-Security"

func SetStrictTransportSecurityHeader(response *http.Response) {
	response.Header.Set(StrictTransportSecurity, "max-age=300; includeSubDomains")
}
