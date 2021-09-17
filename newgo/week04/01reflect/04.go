package main

import (
	"log"
	"reflect"
)

type Person4 struct {
	Name string
	Age  int
}

func (p Person4) ReflectCallFuncWithArge(name string, age int) {
	log.Printf("[调用带参数的方法][arges.name:%v][arges.age:%v]",
		name,
		age,
	)
}

func main() {
	p1 := Person4{
		Name: "leo",
		Age:  12,
	}
	//1.首先通过reflect.ValueOf 获取到反射类型对象
	getValue := reflect.ValueOf(p1)
	//2.获取method对象
	methodValue := getValue.MethodByName("ReflectCallFuncWithArge")
	//造参数
	ages := []reflect.Value{reflect.ValueOf("leo"), reflect.ValueOf(30)}
	methodValue.Call(ages)
}
