package main

import "fmt"

func add1(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}
func calc(a, b int, f func(int, int) int) int {
	return f(a, b)
}
func main() {
	f1 := add1
	fmt.Printf("%T\n", f1)
	ret := calc(100, 200, add1)
	fmt.Println(ret)

	ret = calc(100, 200, sub)
	fmt.Println(ret)
	//匿名函数

	func() {
		fmt.Println("hello world")
	}()
}
