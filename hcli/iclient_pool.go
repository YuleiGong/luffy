package hcli

import (
	"net/http"
	"sync"
)

type IClientPool interface {
	GetOrCreateClientWithTLS(k, caCert string) (*http.Client, error)
	GetOrCreateClient(k string) *http.Client
	DeleteClient(k string)
	IsExist(k string) bool
	ClearClient()
}

var (
	cliPool IClientPool
	once    sync.Once
)

func GetClientPool() IClientPool {
	once.Do(func() {
		cliPool = new(ClientPool)
	})

	return cliPool
}
