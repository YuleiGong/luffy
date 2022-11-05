package hcli

import (
	"net"
)

func IsTimeout(err error) bool {
	oe, ok := err.(net.Error)
	if ok && oe.Timeout() {
		return true
	}
	return false
}
