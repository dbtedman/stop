package http_header

import "net/http"

const XFrameOptions = "X-Frame-Options"

func SetXFrameOptionsHeader(response *http.Response) {
	response.Header.Set(XFrameOptions, "DENY")
}
