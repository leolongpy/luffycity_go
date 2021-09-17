package main

import (
	"log"
	"os"
)

func main() {
	hn, _ := os.Hostname()
	log.Printf("主机名:%v", hn)
	log.Printf("进程pid:%v", os.Getppid())
	log.Printf("命令行参数:%v", os.Args)
	log.Printf("获取GOROOT环境变量:%v", os.Getenv("GOROOT"))

	for _, v := range os.Environ() {
		log.Printf("环境变量 %v", v)
	}
	dir, _ := os.Getwd()
	log.Printf("当前目录：%v", dir)
	//创建单一config目录
	os.Mkdir("config", 0755)
	//创建层级config/yaml/local目录
	os.MkdirAll("config1/yaml/local", 0755)
	//删除单一文件目录
	os.Remove("config")
	//删除层级文件或目录
	os.RemoveAll("config1")

}
