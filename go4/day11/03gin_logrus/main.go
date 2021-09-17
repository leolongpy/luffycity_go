package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var log = logrus.New()

func initLogrus() (err error) {
	// 初始化logrus配置
	log.Formatter = &logrus.JSONFormatter{} // 记录JSON格式的日志
	//置顶日志输出
	file, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Out = file
	// 告诉Gin框架把它的日志也记录到我们打开的文件中
	gin.SetMode(gin.ReleaseMode) //上线的时候设置为ReleaseMode
	gin.DisableConsoleColor()
	gin.DefaultWriter = log.Out
	log.Level = logrus.DebugLevel
	return
}
func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "Hello world",
	})
}

func main() {
	err := initLogrus()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.GET("/index", indexHandler)
	r.Run()
}
