package hcli

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type RequestOpt func(*Request)

type Request struct {
	method  string
	url     string
	timeout time.Duration
	header  map[string]string
	file    uploadFile
}

type uploadFile struct {
	path  string
	field string
	name  string
}

func WithUploadFile(path, field string) RequestOpt {
	return func(r *Request) {
		r.file.path = path
		r.file.field = field
		r.file.name = filepath.Base(path)
	}
}

func WithTimeout(t time.Duration) RequestOpt {
	return func(r *Request) {
		r.timeout = t
	}
}

func WithHeader(h map[string]string) RequestOpt {
	return func(r *Request) {
		for k, v := range h {
			r.header[k] = v
		}
	}
}

func (r *Request) Do(body []byte, cli *http.Client) (resp *http.Response, err error) {
	var req *http.Request
	if req, err = r.wrapRequest(body); err != nil {
		return
	}
	if r.isTimeout() {
		var cancel context.CancelFunc
		cancel, req = r.wrapTimeoutRequest(req)
		defer cancel()
	}

	return cli.Do(req)
}

func (r *Request) isFile() bool {
	return r.file.name != ""
}

func (r *Request) isTimeout() bool {
	return r.timeout != 0
}

func (r *Request) isHeader() bool {
	return len(r.header) != 0
}

func (r *Request) wrapRequest(body []byte) (req *http.Request, err error) {
	if r.isFile() { //文件上传
		req, err = r.wrapFileRequest()
	} else {
		req, err = http.NewRequest(r.method, r.url, bytes.NewReader(body))
	}
	if err != nil {
		return
	}
	if r.isHeader() {
		r.wrapHeaderRequest(req)
	}

	return req, nil
}
func (r *Request) wrapFileRequest() (req *http.Request, err error) {
	fBody := new(bytes.Buffer)
	writer := multipart.NewWriter(fBody)
	var formFile io.Writer
	if formFile, err = writer.CreateFormFile(r.file.field, r.file.name); err != nil {
		return
	}

	var file *os.File
	if file, err = os.Open(r.file.path); err != nil {
		return
	}
	defer file.Close()

	io.Copy(formFile, file)
	r.header[CONTENT_TYPE] = writer.FormDataContentType()
	if err = writer.Close(); err != nil {
		return
	}
	return http.NewRequest(r.method, r.url, fBody)
}

func (r *Request) wrapHeaderRequest(req *http.Request) {
	for k, v := range r.header {
		req.Header.Set(k, v)
	}
}

func (r *Request) wrapTimeoutRequest(req *http.Request) (context.CancelFunc, *http.Request) {
	ctx, cancel := context.WithTimeout(context.TODO(), r.timeout)
	return cancel, req.WithContext(ctx)
}
