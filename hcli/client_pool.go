package hcli

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"sync"
)

var clientPool sync.Map

//添加https认证的ca证书
func newClientWithTLS(caCertPath string) (cli *http.Client, err error) {
	pool := x509.NewCertPool()
	var caCert []byte
	if caCert, err = ioutil.ReadFile(caCertPath); err != nil { //服务的证书
		return
	}
	pool.AppendCertsFromPEM(caCert)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{RootCAs: pool},
	}
	cli = &http.Client{Transport: tr}

	return cli, err
}

func newClient() *http.Client {
	return &http.Client{}
}

func setClient(k string, v *http.Client) {
	clientPool.Store(k, v)
}
