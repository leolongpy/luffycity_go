package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("Hello")
}

func Poni(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("类型：", t)
	fmt.Println("字符串类型:", t.Name())
	//获取值
	v := reflect.ValueOf(o)
	fmt.Println(v)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("%s:%v \n", f.Name, f.Type)

		val := v.Field(i).Interface()
		fmt.Println("val:", val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name)
		fmt.Println(m.Type)
	}
}

func main() {
	u := User{1, "zs", 20}
	Poni(u)

}
