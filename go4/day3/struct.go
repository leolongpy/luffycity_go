package main

import "fmt"

type NewInt int

type haojie = int

func main() {
	var a NewInt
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	var b haojie
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	var c byte
	fmt.Println(c)
	fmt.Printf("%T\n", c)

}
