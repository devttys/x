package http

import (
	"math/rand"
	"net/http"
	"net/url"
)

type ProxyManager struct {
	proxyList []*url.URL
}

func NewProxyManager(urls []*url.URL) *ProxyManager {
	return &ProxyManager{proxyList: urls}
}

func (p *ProxyManager) Proxy(r *http.Request) (*url.URL, error) {
	n := rand.Intn(len(p.proxyList))
	return p.proxyList[n], nil
}
