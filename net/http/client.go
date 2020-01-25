package http

import (
	"net/http"
	"net/url"
	"time"
)

var (
	defaultClient = Client(Transport())
)

func DefaultClient() *http.Client {
	return defaultClient
}

func Transport() *http.Transport {
	return http.DefaultTransport.(*http.Transport).Clone()
}

func WithProxy(proxyFunc func(*http.Request) (*url.URL, error)) *http.Client {
	tr := Transport()

	tr.Proxy = proxyFunc
	return Client(tr)
}

func Client(transport *http.Transport) *http.Client {
	return &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}
}
