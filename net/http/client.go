package http

import (
	"net/http"
	"net/url"
	"time"
)

func DefaultClient() *http.Client {
	return Client(http.DefaultTransport.(*http.Transport))
}

func WithProxy(proxyFunc func(*http.Request) (*url.URL, error)) *http.Client {
	tr := http.DefaultTransport.(*http.Transport).Clone()

	tr.Proxy = proxyFunc
	return Client(tr)
}

func Client(transport *http.Transport) *http.Client {
	return &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}
}
