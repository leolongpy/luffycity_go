package main

import "luffycity_go/go4/day4/mylog"

func main() {
	f1 := mylog.NewFileLogger(mylog.DEBUG, "./", "test.log")
	userID := 10
	f1.Debug("id是%d的用户一直在尝试登陆", userID)
}
