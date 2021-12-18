package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 这是往文件里面写日志的代码

// FileLogger 文件日志结构体
type FileLogger struct {
	level    Level
	fileName string
	filePath string
	file     *os.File
	errFile  *os.File
	maxSize  int64
}

// NewFileLogger 文件日志结构体的构造函数
func NewFileLogger(levelStr, fileName, filePath string) *FileLogger {
	logLevel := paresLogLevel(levelStr)
	fl := &FileLogger{
		level:    logLevel,
		fileName: fileName,
		filePath: filePath,
		maxSize:  10 * 1024 * 1024,
	}
	fl.initFile() // 根据上面的文件路径和文件名打开日志文件，把文件句柄赋值给结构体对应的字段
	return fl
}

// 将指定的日志文件打开 赋值给结构体
func (f *FileLogger) initFile() {
	logName := path.Join(f.filePath, f.fileName)
	//打开文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		panic(fmt.Errorf("打开日志文件%s失败,%v", logName, err))
	}
	f.file = fileObj
	//打开错误的日志文件
	errLogName := fmt.Sprintf("%s.err", logName)
	errfileObj, err := os.OpenFile(errLogName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		panic(fmt.Errorf("打开日志文件%s失败,%v", errLogName, err))
	}
	f.errFile = errfileObj
}

//检查是否要拆分
func (f *FileLogger) checkSplit(file *os.File) bool {
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	return fileSize >= f.maxSize
}

//分装一个切分日志的方法
func (f *FileLogger) splitLogFile(file *os.File) *os.File {
	fileName := file.Name() //文件完整路径
	backupName := fmt.Sprintf("%s_%v.back", fileName, time.Now().Unix())
	//关闭原来的文件
	file.Close()
	//备份原来的文件
	os.Rename(fileName, backupName)
	//新建一个文件
	fileObj, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		panic(fmt.Errorf("打开日志文件失败"))
	}
	return fileObj
}

// 将公用的记录日志的功能封装成一个单独的方法
func (f *FileLogger) log(level Level, format string, args ...interface{}) {
	if f.level > level {
		return
	}
	msg := fmt.Sprintf(format, args...)
	// 日志格式：[时间][文件:行号][函数名][日志级别] 日志信息
	nowStr := time.Now().Format("2006-01-02 15:04:05.000")
	fileName, line, funcName := getCallerInfo(3)
	logLevelStr := getLevelStr(level)
	logMsg := fmt.Sprintf("[%s][%s:%d][%s][%s]%s", nowStr, fileName, line, funcName, logLevelStr, msg)
	if f.checkSplit(f.file) {
		f.file = f.splitLogFile(f.file)
	}
	fmt.Fprintln(f.file, logMsg)
	if level >= ErrorLevel {
		if f.checkSplit(f.errFile) {
			f.errFile = f.splitLogFile(f.errFile)
		}
		fmt.Fprintln(f.errFile, logMsg)
	}
}

// Debug debug方法
func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.log(DebugLevel, format, args...)
}

// Info info方法
func (f *FileLogger) Info(format string, args ...interface{}) {
	f.log(InfoLevel, format, args...)
}

// Warn warn方法
func (f *FileLogger) Warn(format string, args ...interface{}) {
	f.log(WarningLevel, format, args...)
}

// Error error方法
func (f *FileLogger) Error(format string, args ...interface{}) {
	f.log(ErrorLevel, format, args...)
}

// Fatal fatal方法
func (f *FileLogger) Fatal(format string, args ...interface{}) {
	f.log(FatalLevel, format, args...)
}

// Close 关闭日志文件句柄
func (f *FileLogger) Close() {
	f.file.Close()
	f.errFile.Close()
}
