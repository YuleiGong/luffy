# glog
* 用于项目中的日志输出

## 支持
* __WithLogPath__ 自定义日志持久化目录
* __WithLogLevel__ 自定义日志级别
* __WithRotationCount__ 自定义最大日志数，超过会自动删除
* __WithRotationTime__ 自定义日志切分频率
* __WithShowStdOut__ 支持同时输出到StdOut

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



