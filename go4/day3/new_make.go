package main

import "fmt"

func main() {
	var a = new(int)
	fmt.Println(a)

	*a = 10
	fmt.Println(a)
	fmt.Println(*a)

	var c = new([3]int)
	fmt.Println(c)
	c[0] = 1
	fmt.Println(*c)
}
