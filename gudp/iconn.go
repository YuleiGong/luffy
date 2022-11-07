package gudp

import (
	"gudp/handler"
	"gudp/message"
	"gudp/un_pack"
	"net"
)

type IConn interface {
	Do()
	Wait()
}

func NewConn(svr *Server, conn *net.UDPConn, h handler.IHandler, p un_pack.IUnPack) IConn {
	return &Conn{
		conn:        conn,
		svr:         svr,
		handler:     h,
		unPack:      p,
		reveiveChan: make(chan message.Message, svr.reveiveChanSize),
	}
}
