package cmd

import (
	"fmt"
	"net/url"
)

type Options interface {
	ListenAddress() string
	ProxyAddressURL() url.URL
}

var _ Options = theOptions{}

type theOptions struct {
	listenAddress string
	proxyAddress  url.URL
}

func NewOptions(listenAddress string, proxyAddressRaw string) (Options, error) {
	if listenAddress == "" {
		return theOptions{}, fmt.Errorf("the provided listenAddress [%s] is invalid", listenAddress)
	}

	proxyAddress, err := url.Parse(proxyAddressRaw)

	if proxyAddressRaw == "" || err != nil {
		return theOptions{}, fmt.Errorf("the provided proxyAddress [%s] is invalid", proxyAddress)
	}

	return theOptions{
		listenAddress: listenAddress,
		proxyAddress:  *proxyAddress,
	}, nil
}

func (o theOptions) ListenAddress() string {
	return o.listenAddress
}

func (o theOptions) ProxyAddressURL() url.URL {
	return o.proxyAddress
}
