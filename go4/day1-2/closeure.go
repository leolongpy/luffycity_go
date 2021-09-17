package main

import (
	"fmt"
	"strings"
)

func f1(num int) func(int) int {
	f := func(x int) int {
		fmt.Println(num)
		return num + x
	}
	return f
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc1(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
func main() {
	ret := f1(100)
	fmt.Printf("%T\n", ret)
	//fmt.Println(ret(10))
	//fmt.Println(ret(20))
	//fmt.Println(ret(30))
	//fmt.Println(ret(40))

	aviFunc := makeSuffixFunc(".avi")
	fmt.Println(aviFunc("111.avi"))
	textFunc := makeSuffixFunc(".txt")
	fmt.Println(textFunc("豪杰春香"))
	f1, f2 := calc1(100)
	fmt.Println(f1(150), f2(50))

}
