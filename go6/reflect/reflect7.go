package main

import (
	"fmt"
	"reflect"
)

type User3 struct {
	Id   int
	Name string
	Age  int
}

func (u User3) Hello(name string) {
	fmt.Println("Hello:", name)
}
func main() {
	u := User3{1, "zs", 20}
	v := reflect.ValueOf(u)
	//获取方法
	m := v.MethodByName("Hello")
	args := []reflect.Value{reflect.ValueOf("6666")}
	m.Call(args)
}
