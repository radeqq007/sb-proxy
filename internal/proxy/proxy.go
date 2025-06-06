package proxy

import (
	"net/http/httputil"
	"net/url"
)

func newProxy(target string) *httputil.ReverseProxy {
	targetURL, err := url.Parse(target)
	if err != nil {
		panic(err)
	}

	return httputil.NewSingleHostReverseProxy(targetURL)
}
