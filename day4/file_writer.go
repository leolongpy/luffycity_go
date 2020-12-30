package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("2.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	defer file.Close()
	str := "Hello 沙河"
	file.Write([]byte("哈哈哈\n"))
	file.WriteString(str)

}
