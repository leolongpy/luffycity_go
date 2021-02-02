package main

import (
	"fmt"
	"reflect"
)

//通过反射修改值
func modifyValue(x interface{}) {
	v := reflect.ValueOf(x)
	kind := v.Kind()
	fmt.Println(kind)
	if kind == reflect.Ptr {
		// 传入的是一个指针
		// 取到指针对应的值再去修改
		v.Elem().SetInt(1200)
	}
}
func main() {
	var a int64 = 100
	modifyValue(&a)
	fmt.Println(a)
}
