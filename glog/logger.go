package glog

import (
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
)

func newLogger(loggerOpt *LoggerOption) (err error) {
	if err = os.MkdirAll(loggerOpt.logPath, os.ModePerm); err != nil {
		return
	}

	logger = logrus.New()
	logger.SetLevel(logrus.Level(loggerOpt.logLevel))
	if loggerOpt.showStdOut {
		logger.SetOutput(os.Stdout)
	} else {
		logger.SetOutput(ioutil.Discard)
	}

	//设置是否打印调用函数
	logger.SetReportCaller(loggerOpt.reportCaller)
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		FullTimestamp:    true,
		TimestampFormat:  loggerOpt.timestampFormat,
		CallerPrettyfier: prettyCaller,
	})

	return nil

}

func writer(loggerOpt *LoggerOption, level logrus.Level) (writer *rotatelogs.RotateLogs, err error) {
	var logName string
	switch level {
	case logrus.DebugLevel:
		logName = "server.log.%Y-%m-%d"
	case logrus.WarnLevel:
		logName = "warn.log.%Y-%m-%d"
	case logrus.ErrorLevel:
		logName = "error.log.%Y-%m-%d"
	case logrus.PanicLevel:
		logName = "crash.log.%Y-%m-%d"
	default:
		logName = "crash.log.%Y-%m-%d"
	}

	logPath := fmt.Sprintf("%s/%s", loggerOpt.logPath, logName)
	return rotatelogs.New(
		logPath,
		rotatelogs.WithRotationCount(uint(loggerOpt.rotationCount)),        // 文件最大保存份数
		rotatelogs.WithRotationTime(time.Duration(loggerOpt.rotationTime)), // 日志切割时间间隔

	)
}

func InitLogger(loggerOpt *LoggerOption) (err error) {
	if err = newLogger(loggerOpt); err != nil {
		return
	}

	var (
		debugWriter *rotatelogs.RotateLogs
		warnWriter  *rotatelogs.RotateLogs
		errorWriter *rotatelogs.RotateLogs
		crashWriter *rotatelogs.RotateLogs
	)
	if debugWriter, err = writer(loggerOpt, logrus.DebugLevel); err != nil {
		return
	}
	if warnWriter, err = writer(loggerOpt, logrus.WarnLevel); err != nil {
		return
	}
	if errorWriter, err = writer(loggerOpt, logrus.ErrorLevel); err != nil {
		return
	}
	if crashWriter, err = writer(loggerOpt, logrus.PanicLevel); err != nil {
		return
	}

	fileFormatter := &Formatter{
		TimestampFormat: loggerOpt.timestampFormat,
		ShowFullLevel:   loggerOpt.showFullLevel,
		ShowFullPath:    loggerOpt.showFullPath,
		ReportCaller:    loggerOpt.reportCaller,
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.PanicLevel: crashWriter,
		logrus.FatalLevel: crashWriter,
		logrus.ErrorLevel: errorWriter,
		logrus.WarnLevel:  warnWriter,
		logrus.InfoLevel:  debugWriter,
		logrus.DebugLevel: debugWriter,
		logrus.TraceLevel: debugWriter,
	}, fileFormatter)

	logger.AddHook(lfHook)

	//获取崩溃日志文件名
	var crashLogPath string
	crashLogPath = crashWriter.CurrentFileName()
	if crashLogPath == "" {
		year, month, day := time.Now().Date()
		crashLogPath = fmt.Sprintf("%s/crash.log.%04d-%02d-%02d", loggerOpt.logPath, year, month, day)
	}

	fd, err := os.OpenFile(crashLogPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	//stderr 重定向
	redirectStderr(fd)

	return nil
}

//redirectStderr 重定向
func redirectStderr(f *os.File) (err error) {
	err = syscall.Dup2(int(f.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		logger.Errorf("Failed to redirect stderr to file: %v", err)
		return err
	}

	return nil
}
