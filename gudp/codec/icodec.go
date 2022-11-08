package codec

import (
	"gudp/message"
	"net"
)

type ICodec interface {
	Decode(*net.UDPConn) (*message.Message, error)
}

func NewCodec() ICodec {
	return &Codec{}
}
