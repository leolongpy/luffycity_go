package main

import "fmt"

const pi = 3.14

//批量声明变量

//const (
//	a=100
//	b=200
//	c=300
//)

// iota 枚举
//const (
//	aa = iota
//	bb = iota
//	cc = iota
//)

const (
	_  = iota
	KB = 1 << (10 * iota) // 1<<10 2的10次方
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)
const (
	a, b = iota + 1, iota + 2
	c, d = iota + 1, iota + 2
	e, f = iota + 1, iota + 2
)

func main() {
	fmt.Println(pi)
	fmt.Println(a, b, c, d, e, f)
}
