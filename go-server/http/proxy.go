package http

import (
	"net/http/httputil"
	"net/url"
)

func NewProxy(targetAPI string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse("api/py" + targetAPI)
	if err != nil {
		return nil, err
	}
	return httputil.NewSingleHostReverseProxy(url), nil
}
