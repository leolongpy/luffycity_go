package main

import "fmt"

func main() {
	m1 := map[int]string{1: "zs", 2: "ls"}
	//修改
	m1[1] = "ww"
	m1[3] = "zhaosi"
	fmt.Println(m1)

	//创建
	m2 := make(map[int]string, 8)
	m2[0] = "aaaa"
	m2[1] = "bbbb"
	fmt.Println(m2[0])
}
