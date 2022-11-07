package gudp

import (
	"gudp/un_pack"
	"sync"
)

type IServer interface {
	Listen() error
	Start()
}

type ServerOpt func(*Server)

func WithReadBuffSize(size int) ServerOpt {
	return func(s *Server) {
		s.readBuffSiziue = size
	}
}

func WithReceiveChannelSize(size int) ServerOpt {
	return func(s *Server) {
		s.receiveChannelSize = size
	}
}

func NewServer(handler handler.IHanlder, unPack un_pack.IUnPack, opts ...ServerOpt) IServer {
	s := &Server{
		handler:      handler,
		unPack:       unPack,
		readBuffSize: defaultReadBuffSize,
		datagramPool: sync.Pool{
			New: func() interface{} {
				return make([]byte, datagramPoolSize)
			},
		},
	}
	for _, opt := range opts {
		opt(s)
	}

	return s
}
