package main

import "fmt"

func main() {
	var x interface{}
	var a int64 = 100
	var b int32 = 10
	var c int8 = 1
	x = a
	x = b
	x = c
	value, ok := x.(int)
	fmt.Printf("ok:%t value:%#v value type:%T\n", ok, value, value)
}
