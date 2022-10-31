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

func TestSendHTTP(t *testing.T) {
	var (
		r    hcli.IResponse
		err  error
		body []byte
	)
	if r, err = sendHTTP(); err != nil {
		t.Fatal(err)
	}
	t.Logf("%d", r.GetStatusCode())
	t.Logf("%v", r.GetHeader())

	if body, err = r.GetBody(); err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", string(body))

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
	client := cliPool.GetOrCreateClient("http")

	var resp *http.Response
	resp, err = hcli.NewRequest(method, url).Do(hcli.NULL_BODY, client)

	return hcli.NewResponse(resp), err
}
