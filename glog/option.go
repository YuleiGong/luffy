package glog

var (
	defaultLogPath         = "./logs"
	defaultTimestampFormat = "2006-01-02 15:04:05.000"
	defaultLogLevel        = DebugLevel
	defaultReportCaller    = false
	defaultShowFullLevel   = false
	defaultShowFullPath    = false
	defaultShowStdOut      = false
	defaultRotationCount   = 150
	defaultRotationTime    = 86400
)

const (
	PanicLevel int = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

func defaultLoggerOption() *LoggerOption {
	return &LoggerOption{
		logPath:         defaultLogPath,
		logLevel:        defaultLogLevel,
		rotationCount:   defaultRotationCount,
		rotationTime:    defaultRotationTime,
		reportCaller:    defaultReportCaller,
		showFullLevel:   defaultShowFullLevel,
		showFullPath:    defaultShowFullPath,
		showStdOut:      defaultShowStdOut,
		timestampFormat: defaultTimestampFormat,
	}
}

type LoggerOption struct {
	logPath         string
	timestampFormat string
	logLevel        int
	rotationCount   int
	rotationTime    int
	reportCaller    bool
	showFullLevel   bool
	showFullPath    bool
	showStdOut      bool
}
