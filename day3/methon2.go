package main

import "fmt"

type MyInt int

func (m *MyInt) sayHi()  {
	fmt.Println("Hello MyInt~")
}
func main() {
	var a MyInt
	fmt.Println(a)
	a.sayHi()
}
