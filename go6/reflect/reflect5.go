package main

import (
	"fmt"
	"reflect"
)

//定义结构体
type Users struct {
	Id   int
	Name string
	Age  int
}

//匿名字段
type Boy struct {
	Users
	Addr string
}

func main() {
	m := Boy{Users{1, "zs", 20}, "bj"}
	t := reflect.TypeOf(m)
	fmt.Println(t)
	fmt.Printf("%#v\n", t.Field(0))
	fmt.Printf("%#v\n", reflect.ValueOf(m).Field(0))

}
