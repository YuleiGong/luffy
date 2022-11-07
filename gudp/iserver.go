package gudp

type IServer interface {
	SetHandler(IHandler)
	Listen() error
}

func NewServer() IServer {
	return &Server{}
}
