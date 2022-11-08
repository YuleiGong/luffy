package codec

import (
	"gudp/message"
	"net"
)

type Codec struct{}

//数据解包
func (u *Codec) Decode(conn *net.UDPConn) (msg *message.Message, err error) {
	var (
		buf  []byte
		addr net.Addr
	)
	if _, addr, err = conn.ReadFrom(buf); err != nil {
		return
	}

	return message.NewMessage(buf, addr.String()), err
}
