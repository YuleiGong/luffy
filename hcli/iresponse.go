package hcli

import "net/http"

type IResponse interface {
	GetStatusCode() int
	GetHeader() http.Header
	GetBody() ([]byte, error)
	GetStructBody(ptr interface{}) error
	Close()
}

func NewResponse(resp *http.Response) IResponse {
	return &Response{Response: resp}
}
