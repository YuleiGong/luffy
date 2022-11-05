package example

import (
	"hcli"
	"net/http"
	"testing"
	"time"
)

var cliPool hcli.IClientPool = hcli.GetClientPool()

type Resp struct {
	Result string `json:"result"`
}

func TestSendHTTP(t *testing.T) {
	if _, err := sendTimeoutHTTP(); err != nil {
		if hcli.IsTimeout(err) {
			t.Logf("time out")
		} else {
			t.Fatal(err)
		}
	}
}

func sendTimeoutHTTP() (r hcli.IResponse, err error) {
	var (
		method = "GET"
		url    = "http://127.0.0.1:8080/timeout"
	)
	client := cliPool.GetOrCreateClient("http") //获取http client

	var resp *http.Response //发发送一个请求
	resp, err = hcli.NewRequest(method, url, hcli.WithTimeout(1*time.Second)).Do(hcli.NULL_BODY, client)

	return hcli.NewResponse(resp), err //初始化一个Response对象，便于解析结果
}
