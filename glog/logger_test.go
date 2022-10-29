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

func TestWithLogLevel(t *testing.T) {
	var level = map[string]int{
		"Panic": PanicLevel,
		"Fatal": FatalLevel,
		"Error": ErrorLevel,
		"Warn":  WarnLevel,
		"Info":  InfoLevel,
		"Debug": DebugLevel,
		"Trace": TraceLevel,
	}
	for k, lv := range level {
		opts := []LoggerOpt{
			WithShowStdOut(true),
			WithLogLevel(lv),
		}
		t.Logf("log level %s", k)
		if err := InitLogger(NewLoggerOption(opts...)); err != nil {
			t.Errorf("%v", err)
		}
		Info("%v", "test")
		Warn("%v", "warn")
		Debug("%v", "debug")
		Error("%v", "error")
		//Fatal("%v", "fatal")
		//Panic("%v", "panic")
	}
}
