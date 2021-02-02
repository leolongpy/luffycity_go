package main

import (
	"fmt"
	"reflect"
)

// 反射得到结构体方法
type Students struct {
	name  string
	score int
}

func (s Students) Study(day int) string {
	msg := "好好学习，天天向上"
	fmt.Println(msg)
	return msg
}
func (s Students) Sleep() string {
	msg := "好好睡觉，快快长大"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("methon name:%s\n", t.Method(i).Name)
		fmt.Printf("methon:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}
func main() {
	var hj = Students{
		name:  "豪杰",
		score: 90,
	}
	printMethod(hj)
}
