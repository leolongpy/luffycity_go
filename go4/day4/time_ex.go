package main

import (
	"fmt"
	"time"
)

func printTime(t time.Time) {
	timestr := t.Format("2006/01/02 15:04:05")
	fmt.Println(timestr)
}

func calcTime() {
	start := time.Now()
	startTimestamp := start.UnixNano() / 1000
	fmt.Println("《钗头凤》红酥手 黄藤酒 满园春色宫墙柳。")
	time.Sleep(time.Millisecond * 30)
	end := time.Now()
	endTimestamp := end.UnixNano() / 1000
	fmt.Printf("耗费了%d微秒\n", endTimestamp-startTimestamp)
}
func calcTime2() {
	start := time.Now()
	fmt.Println("《钗头凤》红酥手 黄藤酒 满园春色宫墙柳。")
	time.Sleep(time.Millisecond * 30)
	fmt.Println("耗费了", time.Since(start))

}
func main() {
	//now:=time.Now()
	//printTime(now)
	calcTime()
	calcTime2()
}
