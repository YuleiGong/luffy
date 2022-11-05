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

func TestSendHTTP(t *testing.T) {
	var (
		r   hcli.IResponse
		err error
	)
	if r, err = sendHTTP(); err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	resp := Resp{}
	if err = r.GetStructBody(&resp); err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", resp)

}

func sendHTTP() (r hcli.IResponse, err error) {
	var (
		method = "GET"
		url    = "http://127.0.0.1:8080/hello"
	)
	client := cliPool.GetOrCreateClient("http") //获取http client

	var resp *http.Response //发发送一个请求
	resp, err = hcli.NewRequest(method, url).Do(hcli.NULL_BODY, client)

	return hcli.NewResponse(resp), err //初始化一个Response对象，便于解析结果
}
