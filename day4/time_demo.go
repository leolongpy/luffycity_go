package main

import (
	"fmt"
	"time"
)

func timestamp2Timeobj(timeatamp int64) {
	timeObj := time.Unix(timeatamp, 0)
	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%4d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
func tickDemo() {
	ticker := time.Tick(time.Second)

	for i := range ticker {
		fmt.Println(i)
	}
}
func formatDeom() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05.000"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))

}
func main() {
	now := time.Now()
	//fmt.Printf("%v\n", now)
	//fmt.Println(now.Year())
	//fmt.Println(now.Month())
	//fmt.Println(now.Day())
	//fmt.Println(now.Hour())
	//fmt.Println(now.Minute())
	//fmt.Println(now.Second())
	//fmt.Println(now.Nanosecond())

	//fmt.Println(now.Unix())
	//fmt.Println(now.UnixNano())
	tm := now.Unix()
	timestamp2Timeobj(tm)
	//tickDemo()
	formatDeom()
}
