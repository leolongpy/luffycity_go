package main

import (
	"fmt"
	"time"
)

func main() {
	timestamp := time.Now().Unix()
	timeObj := time.Unix(timestamp, 0)
	fmt.Println(timeObj)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	fmt.Printf("%02d-%02d-%02d\n", year, month, day)
	timestr := timeObj.Format("2006-01-02 15:04:05")
	fmt.Println(timestr)
}
