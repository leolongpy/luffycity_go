package main

import (
	"fmt"
	"os"
)

// 写入文件
func main() {
	file, err := os.Create("./aaa.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	for i := 0; i < 5; i++ {
		file.WriteString("ab\n")
		file.Write([]byte("cd\n"))
	}
}
