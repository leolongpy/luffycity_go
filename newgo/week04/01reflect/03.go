package main

import (
	"log"
	"reflect"
)

func main() {
	var num = 3.14
	log.Printf("原始值%v", num)
	pointer := reflect.ValueOf(&num)

	newValue := pointer.Elem()
	newValue.SetFloat(3.15)
	log.Printf("新值%v", num)
}
