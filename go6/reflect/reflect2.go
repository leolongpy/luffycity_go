package main

import (
	"fmt"
	"reflect"
)

func reflect_value(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Println(v)
	k := v.Kind()
	fmt.Println(k)

}

func main() {
	var x float64 = 3.4
	reflect_value(x)
}
