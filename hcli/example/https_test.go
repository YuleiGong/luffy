package example

import (
	"hcli"
	"net/http"
	"testing"
)

//获取一个http client pool
var cliPool hcli.IClientPool = hcli.GetClientPool()

type Resp struct {
	Result string `json:"result"`
}

func TestSendHTTPS(t *testing.T) {
	var (
		r   hcli.IResponse
		err error
	)
	if r, err = sendHTTPS(); err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	t.Logf("%d", r.GetStatusCode())
}

func sendHTTPS() (r hcli.IResponse, err error) {
	var (
		method = "GET"
		url    = "https://localhost:8080/hello"
	)

	var client *http.Client
	if client, err = cliPool.GetOrCreateClientWithTLS("https", "cert/ca.crt"); err != nil {
		return
	}

	var resp *http.Response //发发送一个请求
	resp, err = hcli.NewRequest(method, url).Do(hcli.NULL_BODY, client)

	return hcli.NewResponse(resp), err //初始化一个Response对象，便于解析结果
}
