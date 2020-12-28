package main

import "fmt"

func sayHello()  {
	fmt.Println("hello")
}

func sayHi(name string)  {
	fmt.Printf("Hello %s\n", name)
}

func add(a int,b int) int {
	ret:=a+b
	return ret
}

//可变参数
func add2(a int,b ...int) int {
	ret := a
	fmt.Println(a)
	fmt.Printf("b=%v,type:%T\n",b,b)
	for _,v:=range b{
		ret = ret+v
	}
	return ret
}
func main()  {
	//sayHello()
	//sayHi("leo")
	//println(add(1,1))
	fmt.Println(add2(1,2,3,4,5))
}
