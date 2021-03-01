package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func wr() {
	file, err := os.OpenFile("bbb.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	//获取writer对象
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("Hello\n")
	}
	writer.Flush()
}
func re() {
	file, err := os.Open("bbb.txt")
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		fmt.Println(string(line))
	}
}
func main() {
	re()
}
