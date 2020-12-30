package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./1.txt")
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer file.Close()
	var tmp [128]byte //定义一个字节数组
	for {
		n, err := file.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("文件已经读完了")
			return
		}
		if err != nil {
			fmt.Println("read from file failed, err:", err)
			return
		}
		fmt.Printf("读取了%d个字节\n", n)
		fmt.Println(string(tmp[:]))
	}
}
