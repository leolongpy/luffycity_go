package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type Province struct {
	//id 身份证前两位
	Id       string
	Name     string
	File     *os.File
	chanData chan string
}

//声明等待组
var wg sync.WaitGroup

func main() {
	pMap := make(map[string]*Province)
	ps := []string{"北京市11", "天津市12", "河北省13",
		"山西省14", "内蒙古自治区15", "辽宁省21", "吉林省22",
		"黑龙江省23", "上海市31", "江苏省32", "浙江省33", "安徽省34",
		"福建省35", "江西省36", "山东省37", "河南省41", "湖北省42",
		"湖南省43", "广东省44", "广西壮族自治区45", "海南省46",
		"重庆市50", "四川省51", "贵州省52", "云南省53", "西藏自治区54",
		"陕西省61", "甘肃省62", "青海省63", "宁夏回族自治区64", "新疆维吾尔自治区65",
		"香港特别行政区81", "澳门特别行政区82", "台湾省83"}
	//遍历所有城市，创建实例，省份管道创建
	for _, p := range ps {
		name := p[:len(p)-2]
		id := p[len(p)-2:]
		//创建对象
		province := Province{Id: id, Name: name}
		pMap[id] = &province
		//为当前省份打开一个文件
		file, _ := os.OpenFile("D:/BaiduNetdiskDownload/go6期/省份/"+province.Name+".txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		province.File = file
		defer file.Close()
		province.chanData = make(chan string, 1024)
	}
	//遍历43个省份，开34个文件写数据
	for _, province := range pMap {
		wg.Add(1)
		go writeFile(province)
	}
	//读取优质文件
	file, _ := os.Open("D:/BaiduNetdiskDownload/go6期/kfang_good.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF {
			for _, province := range pMap {
				close(province.chanData)
				fmt.Println(province.Name, "管道已经关闭")
			}
			break
		}
		lineStr := string(lineBytes)
		fileldsSlice := strings.Split(lineStr, ",")
		id := fileldsSlice[1][0:2]
		if province, ok := pMap[id]; ok {
			province.chanData <- (lineStr + "\n")
		} else {
			fmt.Println("未知的省", id)
		}
	}
	wg.Wait()
}

//向文件中写入数据
func writeFile(province *Province) {
	for lineStr := range province.chanData {
		province.File.WriteString(lineStr)
		fmt.Print(province.Name, "写入", lineStr)
	}
	wg.Done()
}
