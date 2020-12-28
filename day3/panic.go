package main

import "fmt"

func f1()  {
	defer func() {
		err := recover()
		fmt.Println("recover抓到了panic异常", err)
	}()
	var a []int
	a[0] = 100
	fmt.Println("panic之后")
}
func main() {
	f1()
	fmt.Println("这是main函数")
}
