package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("xiaolong 123ow 12")
	//每次读取4字节
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				log.Printf("读完了：eof错误：%d", n)
				break
			}
			log.Printf("其他错误:%v", err)
			os.Exit(2)
		}
		log.Printf("[读取到的字节数为:%d][内容:%v]", n, string(p[:n]))
		log.Printf("[读取到的字节数为:%d][内容:%v]", n, string(p))
	}

}
