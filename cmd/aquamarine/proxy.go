package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func NewProxyHandler(toUrl url.URL) *httputil.ReverseProxy {
	proxy := httputil.NewSingleHostReverseProxy(&toUrl)
	proxy.ModifyResponse = func(response *http.Response) error {
		SetStrictTransportSecurityHeader(response)
		SetContentSecurityPolicyHeader(response)
		SetXFrameOptionsHeader(response)
		SetXContentTypeOptionsHeader(response)
		SetReferrerPolicyHeader(response)
		SetPermissionsPolicyHeader(response)

		return nil
	}

	return proxy
}

func ListenHTTPWithProxy(proxyServer Proxy, options Options, errorCh *chan error) {
	aProxy := NewProxyHandler(options.ProxyAddressURL())
	aProxy.Transport = TransportLogger{}

	listenAddress := options.ListenAddress()

	handler := func(writer http.ResponseWriter, request *http.Request) {
		request.Host = request.URL.Host
		aProxy.ServeHTTP(writer, request)
	}

	if err := proxyServer.ListenAndServe(listenAddress, handler); err != nil {
		*errorCh <- err
	} else {
		*errorCh <- nil
	}
}

type TransportLogger struct{}

func (TransportLogger) RoundTrip(req *http.Request) (*http.Response, error) {
	res, err := http.DefaultTransport.RoundTrip(req)

	fmt.Printf("%s %s [%s %s] %s %d\n", time.Now().Format(time.RFC3339), req.RemoteAddr, req.Method, req.RequestURI, res.Status, res.ContentLength)

	return res, err
}
