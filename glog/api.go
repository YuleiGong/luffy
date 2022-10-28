package glog

func Debug(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Info(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warn(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Error(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Fatal(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

type LoggerOpt func(*LoggerOption)

func NewLoggerOption(opts ...LoggerOpt) (l *LoggerOption) {
	l = defaultLoggerOption()
	for _, option := range opts {
		option(l)
	}

	return l
}

//日志目录 默认./logs
func WithLogPath(logPath string) LoggerOpt {
	return func(l *LoggerOption) {
		l.logPath = logPath
	}
}

//日志级别 默认 DebugLevel
func WithLogLevel(logLevel int) LoggerOpt {
	return func(l *LoggerOption) {
		l.logLevel = logLevel
	}
}

//日志最大数量
func WithRotationCount(rotationCount int) LoggerOpt {
	return func(l *LoggerOption) {
		l.rotationCount = rotationCount
	}
}

//日志切分频率
func WithRotationTime(rotationTime int) LoggerOpt {
	return func(l *LoggerOption) {
		l.rotationTime = rotationTime
	}
}

//标准输出
func WithShowStdOut(showStdOut bool) LoggerOpt {
	return func(l *LoggerOption) {
		l.showStdOut = showStdOut
	}
}

//输出函数调用位置
func WithReportCeller(reportCeller bool) LoggerOpt {
	return func(l *LoggerOption) {
		l.reportCaller = reportCeller
	}
}
