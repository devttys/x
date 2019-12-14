package http

import (
	"net"
	"net/http"
	"time"
)

func DefaultTransport() *http.Transport {
	return &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
}

func DefaultClient() *http.Client {
	return &http.Client{
		Timeout:   time.Second * 10,
		Transport: DefaultTransport(),
	}
}
