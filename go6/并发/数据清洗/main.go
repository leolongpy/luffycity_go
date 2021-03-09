package main

import (
	"bufio"
	"github.com/axgle/mahonia"
	"io"
	"os"
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
	file, _ := os.Open("D:/BaiduNetdiskDownload/go6期/kfang.txt")
	defer file.Close()
	//创建优质文件
	goodFile, _ := os.OpenFile("D:/BaiduNetdiskDownload/go6期/kfang_good.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer goodFile.Close()

	badFile, _ := os.OpenFile("D:/BaiduNetdiskDownload/go6期/kfang_bad.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer badFile.Close()

	reader := bufio.NewReader(file)
	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

	}
}
