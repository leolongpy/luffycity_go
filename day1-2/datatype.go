package main

import (
	"fmt"
	"math"
)

func main() {
	var a int = 77
	var b int = 077
	var c int = 0xff
	fmt.Println(a, b, c)
	fmt.Printf("%b\n", a)
	fmt.Printf("%o\n", b)
	fmt.Printf("%x\n", c)
	//c变量的内存地址
	fmt.Printf("%p\n", &c)

	//浮点常量
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	//字符串转义
	fmt.Println("\"C:\\go\"")
	var s1 = "单行文本"
	var s2 = `
多行文本
“不用转义”
`
	fmt.Println(s1)
	fmt.Println(s2)
}
