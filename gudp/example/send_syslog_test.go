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
