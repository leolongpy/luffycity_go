package main

import "fmt"

func main() {
	mapSlice := make([]map[string]int, 3, 10)
	fmt.Println(mapSlice)
	mapSlice = append(mapSlice, map[string]int{"aaa": 10})
	mapSlice = append(mapSlice, map[string]int{"bbb": 100})
	fmt.Println(mapSlice)
	mapSlice[2] = make(map[string]int, 10)
	mapSlice[2]["age"] = 18
	fmt.Println(mapSlice)

	sliceMap := make(map[string][]int, 10)
	sliceMap["haojie"] = make([]int, 3, 10)
	sliceMap["haojie"][0] = 1
	sliceMap["haojie"][1] = 2
	sliceMap["haojie"][2] = 3
	fmt.Println(sliceMap)

}
