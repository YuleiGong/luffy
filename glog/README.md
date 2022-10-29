# glog
* 用于项目中的日志输出

## 支持
* __WithLogPath__ 自定义日志持久化目录
* __WithLogLevel__ 自定义日志级别
* __WithRotationCount__ 自定义最大日志数，超过会自动删除
* __WithRotationTime__ 自定义日志切分频率
* __WithShowStdOut__ 支持同时输出到StdOut
* __WithReportCeller__ 输出函数调用栈

## 快速开始

```go
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
```
## 日志打印级别控制
* 可以灵活调整日志打印级别
```
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
```
