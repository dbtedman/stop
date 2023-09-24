package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dbtedman/stop/aquamarine/internal/options"
	"github.com/dbtedman/stop/aquamarine/internal/proxy"
)

func ListenHTTPWithProxy(proxyServer proxy.Proxy, options options.Options, errorCh *chan error) {
	aProxy := proxy.NewProxyHandler(options.ProxyAddressURL())
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
