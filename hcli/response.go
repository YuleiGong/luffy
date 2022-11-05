package hcli

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	*http.Response
	body []byte
}

func (r *Response) GetStatusCode() int {
	return r.StatusCode
}

func (r *Response) GetHeader() http.Header {
	return r.Header
}

func (r *Response) GetBody() (body []byte, err error) {
	if r.isBody() {
		return r.body, nil
	}

	if r.body, err = ioutil.ReadAll(r.Body); err != nil {
		return
	}

	return r.body, err
}

//传入结构体指针
func (r *Response) GetStructBody(ptr interface{}) (err error) {
	var body []byte
	if body, err = r.GetBody(); err != nil {
		return
	}
	return json.Unmarshal(body, ptr)
}

func (r *Response) isBody() bool {
	return len(r.body) != 0
}

func (r *Response) Close() {
	r.Body.Close()
}
