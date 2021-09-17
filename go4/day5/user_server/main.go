package main

import "luffycity_go/go4/day5/mylogger"

func main() {
	logger := mylogger.NewFileLogger("Info", "./", "xxx.log")
	defer logger.Close()
	//logger := mylogger.NewConsoleLogger("info")
	for {
		logger.Info("Info 这是一条测试的日志")
		logger.Error("Error 这是一条测试的日志")
	}
}
