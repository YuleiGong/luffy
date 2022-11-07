package gudp

import (
	"gudp/handler"
	"gudp/un_pack"
	"net"
	"sync"
)

type Server struct {
	addr            string    //地址
	datagramPool    sync.Pool //byte对象池
	receiveChanSize int       //数据接收channel size
	readBuffSize    int
	handler         handler.IHandler //数据逻辑处理
	unPack          un_pack.IUnPack  //数据解包
}

const (
	defaultReadBuffSize    = 64 * 1024
	defaultReceiveChanSize = 10
	datagramPoolSize       = 65536
)

func (s *Server) listen() (conn *net.UDPConn, err error) {
	var addr *net.UDPAddr
	if addr, err = net.ResolveUDPAddr("udp", s.addr); err != nil {
		return
	}

	if conn, err = net.ListenUDP("udp", addr); err != nil {
		return
	}
	conn.SetReadBuffer(s.readBuffSize)

	return conn, nil
}

func (s *Server) Start() (err error) {
	var conn *net.UDPConn
	if conn, err = s.listen(); err != nil {
		return
	}
	if !s.isHandler() {
		return err
	}

	if !s.isUnPack() {
		return err
	}
	c := NewConn(s, conn, s.handler, s.unPack)
	c.Do()
	c.Wait()
}

func (s *Server) isHandler() bool {
	return s.handler != nil
}

func (s *Server) isUnPack() bool {
	return s.unPack != nil
}
