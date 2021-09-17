package main

import "fmt"

func main() {
	type test struct {
		a int16
		b int16
		c int16
	}
	var t = test{
		a: 1,
		b: 2,
		c: 3,
	}
	fmt.Println(&(t.a))
	fmt.Println(&(t.b))
	fmt.Println(&(t.c))
}
