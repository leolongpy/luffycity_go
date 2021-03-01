package main

import (
	"fmt"
	"io/ioutil"
)

func wrs() {
	err := ioutil.WriteFile("yyy.txt", []byte("hello 你好"), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func res() {
	content, err := ioutil.ReadFile("yyy.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}
func main() {
	wrs()
	res()

}
