package http

import (
	"net/http"
	"time"
)

var (
	DefaultClient = &http.Client{
		Timeout:   time.Second * 10,
		Transport: DefaultTransport,
	}

	DefaultTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
)
