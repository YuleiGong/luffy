package gudp

import (
	"sync"
)

type Server struct {
	datagramPool       sync.Pool    //byte对象池
	receiveChannelSize int          //数据接收channel size
	receiveChannel     chan Message //channel定义
	readBuffSize       int
	wait               sync.WaitGroup
	handler            IHandler //数据封包
	UnPack             IUnPack  //数据解包
	conn               IConn    //udp conn
}

const (
	defaultReadBuffSize       = 64 * 1024
	defaultReceiveChannelSize = 10
	datagramPoolSize          = 65536
)

func (s *Server) Listen() (err error) {
}

func (s *Server) Start() {

}
