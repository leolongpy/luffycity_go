package main

import "fmt"

func main()  {
	var a [5]int
	var b [10]int

	a = [5]int{1,2,3,4,5}
	b = [10]int{1,2,3,4}
	fmt.Println(a)
	fmt.Println(b)

	//var c [3]string = [3]string{"北京","上海","深圳"}
	var c = [3]string{"北京","上海","深圳"}
	//fmt.Println(c)
	var d = [...]int{1,2,546,464,46,464,67,87}
	fmt.Println(d)
	fmt.Printf("c:%T  d:%T\n", c, d)

	//根据索引初始化值
	var e [20]int
	e = [20]int{19:1,18:2}
	fmt.Println(e)

	//遍历数组1
	for i := 0; i<len(a);i++{
		fmt.Println(a[i])
	}
	//for range 循环
	for index,value :=range a{
		fmt.Println(index,a[index],value)
	}
}