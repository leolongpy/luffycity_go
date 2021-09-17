package mylogger

import (
	"fmt"
	"os"
	"time"
)

//往终端打印日志

//终端日志结构体
type ConsoleLogger struct {
	level Level
}

//文件日志结构体的构造函数

func NewConsoleLogger(levelSrt string) *ConsoleLogger {
	logLevel := ParesLogLevel(levelSrt)
	cl := &ConsoleLogger{
		level: logLevel,
	}
	return cl
}

func (c *ConsoleLogger) log(level Level, format string, args ...interface{}) {
	if c.level > level {
		return
	}
	msg := fmt.Sprintf(format, args...)
	nowStr := time.Now().Format("2006-01-02 15:04:06.000")
	fileName, line, funcName := getCallerInfo(3)
	logLevelStr := GetLevelStr(level)
	logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s",
		nowStr, fileName, line, funcName, logLevelStr, msg)
	fmt.Fprintln(os.Stdout, logMsg)

}

// Debug debug方法
func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	c.log(DebugLevel, format, args...)
}

// Info info方法
func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	c.log(InfoLevel, format, args...)
}

// Warn warn方法
func (c *ConsoleLogger) Warn(format string, args ...interface{}) {
	c.log(WarningLevel, format, args...)
}

// Error error方法
func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	c.log(ErrorLevel, format, args...)
}

// Fatal fatal方法
func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	c.log(FatalLevel, format, args...)
}
