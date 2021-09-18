package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"luffycity_go/go4/bloger/controller"
	"luffycity_go/go4/bloger/db"
	"net/http"
	"os"
)

var log = logrus.New()

func main() {
	f, err := os.OpenFile("./log/gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open log file failed,err:%v\n", err)
		return
	}
	log.Out = f
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = log.Out
	log.Level = logrus.DebugLevel
	dsn := "root:root@tcp(127.0.0.1:3306)/go_web?parseTime=true"
	err = db.Init(dsn)
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("views/*")
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "views/404.html", nil)
	})
	r.GET("/", controller.IndexHandler)
	err = r.Run()
	if err != nil {
		panic(err)
	}
}
