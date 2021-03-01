package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	srcFile, err := os.Open("./aaa.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//创建新文件
	dstFile, err2 := os.Create("./aaa2..txt")
	if err2 != nil {
		fmt.Println(err2)
	}
	//缓冲读取
	buf := make([]byte, 1024)
	for {
		n, err := srcFile.Read(buf)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		dstFile.Write(buf[:n])
	}
	srcFile.Close()
	dstFile.Close()

}
