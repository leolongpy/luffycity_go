package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//bufio读取数据
func readByLine() {
	file, err := os.Open("./1.txt")
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	defer file.Close()
	//利用缓冲区从文件读取数据
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Print(str)
			return
		}
		if err != nil {
			fmt.Println("读取文件内容失败")
			return
		}
		fmt.Print(str)
	}
}

//ioutil读取文件
func readFile(filename string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}
func main() {
	readByLine()
	readFile("./1.txt")
}
