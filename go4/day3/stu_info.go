package main

import (
	"fmt"
)

func main() {
	studentMap := make(map[string]map[string]int, 10)
	studentMap["豪杰"] = make(map[string]int, 3)
	studentMap["豪杰"]["id"] = 1
	studentMap["豪杰"]["age"] = 18
	studentMap["豪杰"]["score"] = 90

	studentMap["娜扎"] = make(map[string]int, 3)
	studentMap["娜扎"]["id"] = 2
	studentMap["娜扎"]["age"] = 28
	studentMap["娜扎"]["score"] = 100

	for k, v := range studentMap {
		fmt.Println(k)
		for k2, v2 := range v {
			fmt.Println(k2, v2)
		}
	}

	for k, v := range studentMap {
		id, ok := v["id"]
		if ok {
			if id == 1 {
				fmt.Println("name", k)
				for k2, v2 := range v {
					fmt.Println(k2, v2)
				}
			}
		} else {
			fmt.Println("查无此人")
		}
	}

}
