package glog

import (
	"testing"
)

func TestLog(t *testing.T) {
	o := WithShowStdOut(true)
	if err := InitLogger(NewLoggerOption(o)); err != nil {
		t.Errorf("%v", err)
	}
	Info("%v", "test")
	Error("%v", "error")
	Warn("%v", "warn")
	Debug("%v", "debug")
}
