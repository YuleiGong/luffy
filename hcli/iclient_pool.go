package hcli

import "net/http"

type IClientPool interface {
	GetClientWithTLS(k, caCert string) (*http.Client, error)
	GetClient(k string) *http.Client
	DeleteClient(k string)
	IsExist(k string) bool
	ClearClient()
}

type ClientPool struct{}

var cliPool *ClientPool

func GetCliPool() IClientPool {
	if cliPool == nil {
		cliPool = new(ClientPool)
	}

	return cliPool
}

//根据K从池子中获取client，如果该K没有对应的client,或自动创建并加入池子
func (c *ClientPool) GetClientWithTLS(k, caCert string) (*http.Client, error) {
	v, ok := clientPool.Load(k)
	if !ok {
		cli, err := newClientWithTLS(caCert)
		setClient(k, cli)
		return cli, err
	}

	return v.(*http.Client), nil
}

//根据K从池子中获取client，如果该K没有对应的client,或自动创建并加入池子
func (c *ClientPool) GetClient(k string) *http.Client {
	v, ok := clientPool.Load(k)
	if !ok {
		cli := newClient()
		setClient(k, cli)
		return cli
	}
	return v.(*http.Client)
}

func (c *ClientPool) DeleteClient(k string) {
	clientPool.Delete(k)
}

func (c *ClientPool) IsExist(k string) (ok bool) {
	_, ok = clientPool.Load(k)
	return ok
}

func (c *ClientPool) ClearClient() {
	clientPool.Range(func(k, v interface{}) bool {
		clientPool.Delete(k)
		return true
	})
}
