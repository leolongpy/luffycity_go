package main

import (
	"fmt"
)

func main()  {
	var a = 100
	fmt.Printf("a=%v\n",a)
	fmt.Printf("a的类型是%T\n",a)

	fmt.Printf("|%d|\n", a)
	fmt.Printf("|%8d|\n", a)
	fmt.Printf("|%-8d|\n", a)
	fmt.Printf("|%08d|\n", a)

	f1 := 3.141592654
	fmt.Printf("%.2f\n", f1)
	fmt.Printf("%.2g\n", f1)

	//字符串
	s1:="这是一个字符串\""
	fmt.Printf("s1:%s\n", s1)
	fmt.Printf("s1:%q\n", s1)

	fmt.Printf("s1:%50s\n", s1)
	fmt.Printf("s1:%.5s\n", s1)
}
