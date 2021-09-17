package mylog

import (
	"fmt"
	"os"
	"time"
)

//往文件记录日志的结构体
type FileLogger struct {
	level       int //只有大于这个级别的日志才记录
	logFilePath string
	logFileName string
	logFile     *os.File
}

// 是一个生成文件日志结构体实例的构造函数
func NewFileLogger(level int, logFilePath, logFileName string) *FileLogger {
	flObj := &FileLogger{
		level:       level,
		logFileName: logFileName,
		logFilePath: logFilePath,
	}
	flObj.inintFileLogger() // 调用下面的初始化文件句柄的方法
	return flObj
}

// 专门用来初始化文件日志的文件句柄
func (f *FileLogger) inintFileLogger() {
	// 打开日志文件
	filepath := fmt.Sprintf("%s/%s", f.logFilePath, f.logFileName)
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Sprintf("open file:%s failed", filepath))
	}
	f.logFile = file
}

//记录日志
func (f *FileLogger) Debug(format string, args ...interface{}) {
	if f.level > DEBUG {
		return
	}
	fileName, funcName, line := getCallerInfo()
	nowStr := time.Now().Format("[2006-01-02 15:04:05.000]")
	format = fmt.Sprintf("%s [%s] [%s:%s] [%d] %s", nowStr, getLevelStr(f.level), fileName, funcName, line, format)
	fmt.Fprintf(f.logFile, format, args...)
	fmt.Fprintln(f.logFile) //加换行
}

func (f *FileLogger) Info(msg string) {
	f.logFile.WriteString(msg)
}

func (f *FileLogger) Error(msg string) {
	f.logFile.WriteString(msg)
}
