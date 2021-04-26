package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/axgle/mahonia"
)

// 处理乱码
// 参数1：处理的数据
// 参数2：数据目前的编码
// 参数3：返回的正常数据
func ConvertEncoding(srcStr string, encoding string) (dstStr string) {
	enc := mahonia.NewDecoder(encoding)
	utfStr := enc.ConvertString(srcStr)
	dstStr = utfStr
	return
}

func read() {
	contentBytes, err := ioutil.ReadFile("D:/BaiduNetdiskDownload/go6期/kfang.txt")
	if err != nil {
		fmt.Println("读取失败，", err)
	}
	contentStr := string(contentBytes)
	lineStrs := strings.Split(contentStr, "\n\r")
	for _, lineStr := range lineStrs {
		newStr := ConvertEncoding(lineStr, "GBK")
		fmt.Println(newStr)
	}
}

//缓冲读取
func read2() {
	file, _ := os.Open("D:/BaiduNetdiskDownload/go6期/kfang.txt")
	defer file.Close()
	//建缓冲区
	reader := bufio.NewReader(file)
	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		gbkStr := string(lineBytes)
		utfStr := ConvertEncoding(gbkStr, "GBK")
		fmt.Println(utfStr)
	}
}

func main() {
	read2()
}
