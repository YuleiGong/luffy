package glog

import (
	"testing"
)

func TestLog(t *testing.T) {
	o := WithShowStdOut(true)
	if err := InitLogger(NewLoggerOption(o)); err != nil {
		t.Errorf("%v", err)
	}
	DefaultGLog.Info("%v", "test")
	DefaultGLog.Error("%v", "error")
	DefaultGLog.Warn("%v", "warn")
	DefaultGLog.Debug("%v", "debug")
}
