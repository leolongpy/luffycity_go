package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 方式一
	bytes1, _ := ioutil.ReadFile("01.go")
	file, _ := os.Open("01.go")
	// 方式二
	bytes2, _ := ioutil.ReadAll(file)
	//方法三
	file.Close()
	file, _ = os.Open("01.go")
	bo := bufio.NewReader(file)
	buf := make([]byte, 20000)
	bo.Read(buf)

	fmt.Println(string(bytes1))
	fmt.Println(string(bytes2))
	fmt.Println(string(buf))
}
