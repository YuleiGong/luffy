package hcli

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"sync"
)

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

type ClientPool struct {
	pool sync.Map
}

func (c *ClientPool) setClient(k string, v *http.Client) {
	c.pool.Store(k, v)
}

//根据K从池子中获取client，如果该K没有对应的client,或自动创建并加入池子
func (c *ClientPool) GetOrCreateClientWithTLS(k, caCert string) (*http.Client, error) {
	v, ok := c.pool.Load(k)
	if !ok {
		cli, err := newClientWithTLS(caCert)
		c.setClient(k, cli)
		return cli, err
	}

	return v.(*http.Client), nil
}

//根据K从池子中获取client，如果该K没有对应的client,或自动创建并加入池子
func (c *ClientPool) GetOrCreateClient(k string) *http.Client {
	v, ok := c.pool.Load(k)
	if !ok {
		cli := newClient()
		c.setClient(k, cli)
		return cli
	}
	return v.(*http.Client)
}

func (c *ClientPool) DeleteClient(k string) {
	c.pool.Delete(k)
}

func (c *ClientPool) IsExist(k string) (ok bool) {
	_, ok = c.pool.Load(k)
	return ok
}

func (c *ClientPool) ClearClient() {
	c.pool.Range(func(k, v interface{}) bool {
		c.pool.Delete(k)
		return true
	})
}
