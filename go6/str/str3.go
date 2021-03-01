package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	a := strconv.FormatBool(false)
	b := strconv.FormatInt(-1234, 10)
	c := strconv.FormatUint(1234, 10)
	d := strconv.Itoa(-2234)
	fmt.Println(a, b, c, d)
	fmt.Println(reflect.TypeOf(b))
}
