package gudp

import (
	"gudp/handler"
	"gudp/un_pack"
	"sync"
)

type IServer interface {
	Start() error
}

type ServerOpt func(*Server)

func WithReadBuffSize(size int) ServerOpt {
	return func(s *Server) {
		s.readBuffSize = size
	}
}

func WithReceiveChannelSize(size int) ServerOpt {
	return func(s *Server) {
		s.receiveChannelSize = size
	}
}

func NewServer(h handler.IHandler, p un_pack.IUnPack, opts ...ServerOpt) IServer {
	s := &Server{
		handler:      h,
		unPack:       p,
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
