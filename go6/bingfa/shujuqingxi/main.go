package main

import (
	"bufio"
	"fmt"
	"github.com/axgle/mahonia"
	"io"
	"os"
	"strings"
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
func main() {
	//打开文件
	file, _ := os.Open("E:/Desktop/kfang.txt")
	defer file.Close()
	//创建优质文件
	goodFile, _ := os.OpenFile("E:/Desktop/kfang_good.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer goodFile.Close()

	badFile, _ := os.OpenFile("E:/Desktop/kfang_bad.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer badFile.Close()

	reader := bufio.NewReader(file)
	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		gbkStr := string(lineBytes)
		lineStr := ConvertEncoding(gbkStr, "GBK")
		//取身份证
		fields := strings.Split(lineStr, ",")
		if len(fields) >= 2 && len(fields[1]) == 18 {
			goodFile.WriteString(lineStr + "\n")
			fmt.Println("Good:", lineStr)
		} else {
			badFile.WriteString(lineStr + "\n")
			fmt.Println("Bad:", lineStr)
		}

	}
}
