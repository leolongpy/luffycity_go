package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Ager struct {
	decade   string
	file     *os.File
	chanData chan string
}

var wg sync.WaitGroup

func writeFile(ager *Ager) {
	for contentStr := range ager.chanData {
		ager.file.WriteString(contentStr)
		fmt.Println(ager.decade, "-----", contentStr)
	}
	wg.Done()
}

func main() {
	agerMap := make(map[string]*Ager)
	for i := 190; i < 202; i++ {
		ager := Ager{decade: strconv.Itoa(i)}
		file, _ := os.OpenFile("D:/BaiduNetdiskDownload/go6期/年龄/"+ager.decade+".txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		ager.file = file
		defer ager.file.Close()
		ager.chanData = make(chan string)
		agerMap[ager.decade] = &ager
	}
	for _, ager := range agerMap {
		wg.Add(1)
		go writeFile(ager)
	}
	//读取优质文件
	file, _ := os.Open("D:/BaiduNetdiskDownload/go6期/kfang_good.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		linstr, err := reader.ReadString('\n')
		if err == io.EOF {
			for _, ager := range agerMap {
				close(ager.chanData)
			}
			break
		}
		decade := strings.Split(linstr, ",")[1][6:9]

		if ager := agerMap[decade]; ager != nil {
			agerMap[decade].chanData <- linstr + "\n"
		} else {
			fmt.Println(linstr)
		}
	}
	wg.Wait()
}
