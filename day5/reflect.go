package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", t)
	fmt.Printf("%T\n", x)
	fmt.Printf("name:%v kind:%v\n", t.Name(), t.Kind())
}

type cat struct {
	name string
}

type person struct {
	name string
	age  uint8
}

func main() {
	//reflectType(100)
	//reflectType(false)
	//reflectType("leo")
	//reflectType([3]int{1,2,3})
	//reflectType(map[string]int{})

	//自定义结构体
	//var c1 = cat{
	//	name: "花花",
	//}
	//
	//var p1 = person{
	//	name: "豪杰",
	//	age: 18,
	//}
	//
	//reflectType(c1)
	//reflectType(p1)

	//var i int32 = 100
	//var f float32 = 12.34
	//reflectType(&i)
	//reflectType(&f)

	var a = []int{1, 2, 3}
	reflectType(a)
	var b = [3]int{1, 2, 3}
	reflectType(b)
	reflectType(map[string]int{})
}
