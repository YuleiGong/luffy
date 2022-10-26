# 日志模块
* 用于项目中的日志输出，日志文件会定向到指定目录。支持日志切分，日志自动删除。

# 快速开始

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



