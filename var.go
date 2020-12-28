package main

import "fmt"

//全局变量
//var a = 10

func testGloble() {
	a := 100
	b := 200
	fmt.Println(a)
	fmt.Println(b)
}
func main() {
	testGloble()
	fmt.Println(a)

}
