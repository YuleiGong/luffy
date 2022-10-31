package hcli

import (
	"net/http"
	"strings"
)

type IRequest interface {
	Do([]byte, *http.Client) (*http.Response, error)
}

func NewRequest(method, url string, opts ...RequestOpt) IRequest {
	r := &Request{
		method: strings.ToUpper(method),
		url:    url,
		header: make(map[string]string),
	}

	for _, opt := range opts {
		opt(r)
	}

	return r
}
