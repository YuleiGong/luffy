package gudp

import (
	"gudp/handler"
	"gudp/message"
	"gudp/un_pack"
	"net"
	"sync"
	"time"
)

type Conn struct {
	svr         *Server
	conn        *net.UDPConn
	receiveChan chan message.Message //channel定义
	wait        sync.WaitGroup
	handler     handler.IHandler //数据逻辑处理
	unPack      un_pack.IUnPack  //数据解包
}

//reveive
//handler
func (c *Conn) Do() {

}

func (c *Conn) close() {
}

func (c *Conn) receiveMessage() {
	c.wait.Add(1)
	go func() {
		defer c.wait.Done()
		for {
			buf := c.svr.datagramPool.Get().([]byte)
			n, addr, err := c.conn.ReadFrom(buf)
			if err != nil {
				opError, ok := err.(*net.OpError)
				if (ok) && !opError.Temporary() && !opError.Timeout() {
					return
				}
				time.Sleep(10 * time.Millisecond)
			}
			//解包 放入receiveMessage
		}
	}()
}

func (c *Conn) handler() {
	c.wait.Add(1)
	go func() {
		defer c.wait.Done()
		for {
			select {
			case msg, ok := (<-c.receiveMessage):
				if !ok {
					return
				}
				go c.handler.Do(msg) //执行业务处理函数
				c.svr.datagramPool.Put(msg.message[:cap(msg.message)])
			}
		}
	}()
}

func (c *Conn) Wait() {
	c.wait.Wait()
}
