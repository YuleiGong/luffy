package glog

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	logrusPackage      string
	minimumCallerDepth int
	callerInitOnce     sync.Once
)

const (
	maximumCallerDepth = 25
	knownLogrusFrames  = 12
)

type Formatter struct {
	TimestampFormat string
	ShowFullLevel   bool
	ShowFullPath    bool
	ReportCaller    bool
}

func init() {
	minimumCallerDepth = 1
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = time.StampMilli
	}

	// output buffer
	b := &bytes.Buffer{}

	// write time
	b.WriteString("[")
	b.WriteString(entry.Time.Format(timestampFormat))
	b.WriteString("]")

	// write level
	level := strings.ToUpper(entry.Level.String())

	b.WriteString(" [")
	if f.ShowFullLevel {
		b.WriteString(level)
	} else {
		b.WriteString(level[:4])
	}
	b.WriteString("]")

	if f.ReportCaller {
		entry.Caller = setCaller()
		f.writeCaller(b, entry)
	}

	b.WriteString(" - ")
	b.WriteString(strings.TrimSpace(entry.Message))
	b.WriteByte('\n')

	return b.Bytes(), nil
}

func (f *Formatter) writeCaller(b *bytes.Buffer, entry *logrus.Entry) {
	if entry.HasCaller() {
		//是否显示绝对路径
		var filePath string
		if f.ShowFullPath {
			filePath = entry.Caller.File
		} else {
			filePath = getLastFileName(entry.Caller.File)
		}

		fmt.Fprintf(
			b,
			" [%s:%d %s()]",
			filePath,
			entry.Caller.Line,
			getLastFileName(entry.Caller.Function),
		)
	}
}

func getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
}

func getLastFileName(f string) string {
	lastSlash := strings.LastIndex(f, "/")
	if lastSlash != -1 {
		return f[lastSlash+1:]
	}

	return f
}

func setCaller() *runtime.Frame {
	callerInitOnce.Do(func() {
		pcs := make([]uintptr, maximumCallerDepth)
		_ = runtime.Callers(0, pcs)

		for i := 0; i < maximumCallerDepth; i++ {
			funcName := runtime.FuncForPC(pcs[i]).Name()
			if strings.Contains(funcName, "setCaller") {
				logrusPackage = getPackageName(funcName)
				break
			}
		}

		minimumCallerDepth = knownLogrusFrames
	})

	pcs := make([]uintptr, maximumCallerDepth)
	depth := runtime.Callers(minimumCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	for f, again := frames.Next(); again; f, again = frames.Next() {
		pkg := getPackageName(f.Function)
		if pkg != logrusPackage {
			return &f
		}
	}

	return nil
}

func prettyCaller(caller *runtime.Frame) (function, file string) {
	return fmt.Sprintf("%s()] -", getLastFileName(caller.Function)), fmt.Sprintf(" [%s:%d ", getLastFileName(caller.File), caller.Line)
}
