package main

import "fmt"

type Cat1 struct{}

func ShowType(x interface{}) {
	v1, ok := x.(int)
	if !ok {
		fmt.Println("不是int")
	} else {
		fmt.Println("x就是一个int类型", v1)
	}
	v2, ok := x.(string)
	if !ok {
		fmt.Println("不是string")
	} else {
		fmt.Println("x是一个string类型", v2)
	}
}
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string,value is %v\n", v)
	case int:
		fmt.Printf("x is a int,value  is %v\n", v)
	case bool:
		fmt.Printf("x is a bool, value is %v\n", v)
	case Cat1:
		fmt.Printf("x is a Cat struct,value is %v\n", v)
	case *string:
		fmt.Printf("x is a string poninter,value is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
func main() {
	var x interface{}
	x = 100
	ShowType(x)
	ShowType("哈哈")
	justifyType(100)
	justifyType(Cat1{})
	s := "哈哈"
	justifyType(&s)

}
