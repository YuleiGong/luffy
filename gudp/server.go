package gudp

import (
	"fmt"
	"gudp/codec"
	"gudp/handler"
	"net"
)

type Server struct {
	addr            string //地址
	receiveChanSize int    //数据接收channel size
	readBuffSize    int
	stopped         bool
	codec           codec.ICodec     //数据编解码
	handler         handler.IHandler //数据的处理逻辑
}

const (
	defaultReadBuffSize    = 64 * 1024
	defaultSendChanSize    = 10
	defaultReceiveChanSize = 10
)

type ServerOpt func(*Server)

func WithReadBuffSize(size int) ServerOpt {
	return func(s *Server) {
		s.readBuffSize = size
	}
}

func WithReceiveChanSize(size int) ServerOpt {
	return func(s *Server) {
		s.receiveChanSize = size
	}
}

func NewServer(addr string, opts ...ServerOpt) *Server {
	s := &Server{
		addr:            addr,
		receiveChanSize: defaultReceiveChanSize,
		readBuffSize:    defaultReadBuffSize,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Server) RegisterHandler(h handler.IHandler) {
	s.handler = h
}

func (s *Server) RegisterCodec(c codec.ICodec) {
	s.codec = c
}

func (s *Server) Start() (err error) {
	var conn *net.UDPConn
	if conn, err = s.listen(); err != nil {
		return
	}
	if !s.isHandler() {
		return err
	}

	if !s.isCodec() {
		return err
	}
	c := NewConn(s, conn)
	c.Do()
	c.Wait()
	defer conn.Close()

	return nil
}

func (s *Server) listen() (conn *net.UDPConn, err error) {
	var addr *net.UDPAddr
	if addr, err = net.ResolveUDPAddr("udp", s.addr); err != nil {
		return
	}

	if conn, err = net.ListenUDP("udp", addr); err != nil {
		return
	}
	if err = conn.SetReadBuffer(s.readBuffSize); err != nil {
		return
	}

	return conn, nil
}

func (s *Server) isHandler() bool {
	return s.handler != nil
}

func (s *Server) isCodec() bool {
	return s.codec != nil
}

func (s *Server) GetHandler() handler.IHandler {
	return s.handler
}

func (s *Server) GetCodec() codec.ICodec {
	return s.codec
}

func (s *Server) Stop() {
	fmt.Println("**************************")
	s.stopped = true
}

func (s *Server) isStopped() bool {
	return s.stopped
}
