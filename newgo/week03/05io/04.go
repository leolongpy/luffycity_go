package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, err := ioutil.ReadFile("01.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", bytes)
	fmt.Printf("%v", string(bytes))
}
