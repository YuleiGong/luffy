package hcli

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	*http.Response
}

func NewResponse(resp *http.Response) IResponse {
	return &Response{resp}
}

func (r *Response) GetStatusCode() int {
	return r.StatusCode
}

func (r *Response) GetHeader() http.Header {
	return r.Header
}

func (r *Response) GetBody() (body []byte, err error) {
	return ioutil.ReadAll(r.Body)
}

//传入结构体指针
func (r *Response) GetStructBody(ptr interface{}) (err error) {
	var body []byte
	if body, err = r.GetBody(); err != nil {
		return
	}
	return json.Unmarshal(body, ptr)
}
