<a name="d7127e45"></a>
# udp 服务端

- 支持接口形式的拆包和解包，handler处理。
- 基于单向数据流。实际使用中，udp协议不可靠，双向数据流没有意义。
<a name="c5b127aa"></a>
# 功能模块

- Server： 实现了udp conn 的和其他配置的封装，提供了拆包和解包接口，数据处理handler接口。
- Conn: 实现了udp conn的封装，并关联了IServer,调用数据的解包，封包，数据handler处理。
- IHandler: 实现了数据handler接口，在构建udp服务时，使用者需要自己实现Handler接口。
- ICodec: 实现了数据拆包接口，在构建udp服务时，使用者需要自己实现Codec接口。
<a name="BDIzb"></a>
# 一个UDP数据报的处理流程
![](https://cdn.nlark.com/yuque/0/2022/jpeg/2172986/1668857194329-95acda36-a71d-4d31-9e63-5ac1ef6e88fa.jpeg)
<a name="dOB2b"></a>
# 快速使用

- 启动一个UDP服务，自定义decode 和 handler，详细示例见 **example/server.go**
```go
package main

import (
	"fmt"
	"gudp"
	"gudp/message"
	"net"
	"os"
	"os/signal"
	"syscall"
)
//decode接口
type LogCodec struct{}

func (l LogCodec) Decode(conn *net.UDPConn) (msg *message.Message, err error) {
	var addr net.Addr
	buf := make([]byte, 65536)
	if _, addr, err = conn.ReadFrom(buf); err != nil {
		return
	}

	return message.NewMessage(buf, addr.String()), nil
}
//数据处理handler
type LogHandler struct{}

func (l LogHandler) Do(msg *message.Message) {
	fmt.Printf("%s \n", string(msg.GetMessage()))
	fmt.Printf("%s \n", msg.GetClient())
}

func main() {
	addr := "127.0.0.1:8899"
	s := gudp.NewServer(addr)
	s.RegisterHandler(LogHandler{})
	s.RegisterCodec(LogCodec{})

	go func() {
		if err := s.Start(); err != nil {
			fmt.Println(err)
			return
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	s.Stop()
}

```

- 测试代码：
```go
package main

import (
	"net"
	"testing"
	"time"
)

var exampleRFC5424Syslog = "<34>1 2003-10-11T22:14:15.003Z mymachine.example.com su - ID47 - 'su root' failed for lonvick on /dev/pts/8"

func TestSendUDPSyslog(t *testing.T) {
	serverAddr, _ := net.ResolveUDPAddr("udp", "localhost:8899")
	con, _ := net.DialUDP("udp", nil, serverAddr)
	for i := 0; i < 10000; i++ {
		time.Sleep(time.Second)
		if _, err := con.Write([]byte(exampleRFC5424Syslog)); err != nil {
			t.Fatalf("%v", err)
		}
	}
}
```
