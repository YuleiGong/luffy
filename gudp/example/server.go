package main

import (
	"fmt"
	"gudp"
	"gudp/message"
	"net"
)

type LogCodec struct{}

func (l LogCodec) Decode(conn *net.UDPConn) (msg *message.Message, err error) {
	var addr net.Addr
	buf := make([]byte, 65536)
	if _, addr, err = conn.ReadFrom(buf); err != nil {
		return
	}

	return message.NewMessage(buf, addr.String()), nil
}

type LogHandler struct{}

func (l LogHandler) Do(msg *message.Message) {
	fmt.Printf("%s \n", string(msg.GetMessage()))
	fmt.Printf("%s \n", msg.GetClient())
}

func main() {
	addr := "127.0.0.1:8899"
	s := gudp.NewServer(addr)
	s.RegisterHandler(LogHandler{})
	s.RegisterCodec(LogCodec{})
	if err := s.Start(); err != nil {
		fmt.Println(err)
	}
}
