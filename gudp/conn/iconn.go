package conn

type IConn interface{}

func NewConn() IConn {

	return &Conn{}
}
