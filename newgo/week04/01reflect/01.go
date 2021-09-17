package main

import (
	"log"
	"reflect"
)

func main() {
	var name = "leo"
	//TypeOf  会返回目标的类型
	reflectType := reflect.TypeOf(name)
	//返回值
	reflectValue := reflect.ValueOf(name)

	log.Printf("[typeof:%v]", reflectType)
	log.Printf("[valueof:%v]", reflectValue)

}
