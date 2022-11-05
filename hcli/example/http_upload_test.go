package example

import (
	"hcli"
	"net/http"
	"testing"
)

var cliPool hcli.IClientPool = hcli.GetClientPool()

type Resp struct {
	Result string `json:"result"`
}

func TestUpload(t *testing.T) {
	var (
		r    hcli.IResponse
		err  error
		body []byte
	)
	if r, err = sendHTTP(); err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	t.Logf("%d", r.GetStatusCode())

	if body, err = r.GetBody(); err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", string(body))

}

func sendHTTP() (r hcli.IResponse, err error) {
	var (
		method = "GET"
		url    = "http://127.0.0.1:8080/upload"
	)

	client := cliPool.GetOrCreateClient("upload")
	file := "upload.txt"
	field := "upload"

	var resp *http.Response //发发送一个请求
	resp, err = hcli.NewRequest(method, url, hcli.WithUploadFile(file, field)).Do(hcli.NULL_BODY, client)

	return hcli.NewResponse(resp), err //初始化一个Response对象，便于解析结果
}
