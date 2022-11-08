package main

import (
	"net"
	"sync"
	"testing"
)

var exampleRFC5424Syslog = "<34>1 2003-10-11T22:14:15.003Z mymachine.example.com su - ID47 - 'su root' failed for lonvick on /dev/pts/8"

func TestSendUDPSyslog(t *testing.T) {
	wg := sync.WaitGroup{}
	count := 20

	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			serverAddr, _ := net.ResolveUDPAddr("udp", "localhost:8899")
			con, _ := net.DialUDP("udp", nil, serverAddr)
			for i := 0; i < 10000; i++ {
				con.Write([]byte(exampleRFC5424Syslog))
			}
			defer wg.Done()
		}()
	}
	wg.Wait()
}
