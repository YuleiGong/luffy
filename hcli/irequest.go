package hcli

import (
	"net/http"
)

type IRequest interface {
	Do([]byte, *http.Client) (*http.Response, error)
}
