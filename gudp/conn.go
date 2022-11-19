package gudp

import (
	"fmt"
	"gudp/message"
	"net"
	"sync"
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

func (c *Conn) Do() {
	c.receiveMessage()
	c.handler()
}

func (c *Conn) receiveMessage() {
	c.wait.Add(1)
	go func() {
		defer c.wait.Done()
		for {
			if c.svr.isStopped() {
				fmt.Println("stop receiveMessage")
				return
			}
			msg, err := c.svr.GetCodec().Decode(c.conn)
			if err != nil {
				opError, ok := err.(*net.OpError)
				if (ok) && !opError.Temporary() && !opError.Timeout() {
					fmt.Printf("err %v \n", err)
					return
				}
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
			if c.svr.isStopped() {
				fmt.Println("stop handler")
				return
			}
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
