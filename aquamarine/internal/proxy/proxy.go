package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/dbtedman/stop/aquamarine/internal/security"
)

func NewProxyHandler(toUrl url.URL) *httputil.ReverseProxy {
	proxy := httputil.NewSingleHostReverseProxy(&toUrl)
	proxy.ModifyResponse = func(response *http.Response) error {
		security.SetStrictTransportSecurityHeader(response)
		security.SetContentSecurityPolicyHeader(response)
		security.SetXFrameOptionsHeader(response)
		security.SetXContentTypeOptionsHeader(response)
		security.SetReferrerPolicyHeader(response)
		security.SetPermissionsPolicyHeader(response)

		return nil
	}

	return proxy
}
