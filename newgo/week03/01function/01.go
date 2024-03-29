package main

import "fmt"

func f5() (s1 []int) {
	s1 = []int{1, 1}
	defer func() {
		s1[1] = 10
	}()
	return []int{3, 3}
}

/*
1.先给返回值赋值
2.把defer改造成正常函数
3.空的return
*/
func f55() (s1 []int) {
	s1 = []int{1, 1}
	s1 = []int{3, 3}
	func() {
		s1[1] = 10 //闭包
	}()
	return
}
func main() {
	fmt.Println(f5())
	fmt.Println(f55())
}
