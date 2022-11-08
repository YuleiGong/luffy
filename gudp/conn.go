package gudp

import (
	"gudp/message"
	"net"
	"sync"
	"time"
)

type Conn struct {
	svr         *Server
	conn        *net.UDPConn
	receiveChan chan *message.Message //channel定义
	wait        sync.WaitGroup
}

func NewConn(svr *Server, conn *net.UDPConn) *Conn {
	return &Conn{
		svr:         svr,
		conn:        conn,
		receiveChan: make(chan *message.Message, svr.receiveChanSize),
	}
}

//reveive
//handler
func (c *Conn) Do() {
	c.receiveMessage()
	c.handler()

}

//接收消息 解包
func (c *Conn) receiveMessage() {
	c.wait.Add(1)
	go func() {
		defer c.wait.Done()
		for {
			msg, err := c.svr.GetCodec().Decode(c.conn)
			if err != nil {
				opError, ok := err.(*net.OpError)
				if (ok) && !opError.Temporary() && !opError.Timeout() {
					return
				}
				time.Sleep(10 * time.Millisecond)
			}
			c.receiveChan <- msg
		}
	}()
}

//消息handler
func (c *Conn) handler() {
	c.wait.Add(1)
	go func() {
		defer c.wait.Done()
		for {
			select {
			case msg, ok := (<-c.receiveChan):
				if !ok {
					return
				}
				go c.svr.GetHandler().Do(msg) //执行业务处理函数
			}
		}
	}()
}

func (c *Conn) Wait() {
	c.wait.Wait()
}
