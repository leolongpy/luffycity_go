package main

import "github.com/sirupsen/logrus"

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)
	logrus.WithFields(logrus.Fields{
		"name": "leo",
		"age":  90,
	}).Warn("这个是一条warning日志")
	logrus.Infof("这是一条普通日志")
}
