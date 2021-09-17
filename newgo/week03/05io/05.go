package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filename := "a.txt"
	err := ioutil.WriteFile(filename, []byte("升值\n加薪"), 0644)
	fmt.Println(err)
}
