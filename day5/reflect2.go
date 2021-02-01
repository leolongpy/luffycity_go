package main

import (
	"reflect"
)

func reflectValue(x interface{}) {
	v := reflect.TypeOf(x)
	k := v.Kind()
	//fmt.Println(v.Float())
	switch k {
	case reflect.Int64:
		//fmt.Printf("type is int64, value is %d\n",int64(v.Int()))

	}
}
func main() {
	reflectValue(100)
}
