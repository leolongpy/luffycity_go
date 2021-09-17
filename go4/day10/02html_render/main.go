package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func loginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"msg": "嘿嘿",
	})
}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"msg": "哈哈",
	})
}

func main() {
	r := gin.Default()
	//加载模板文件
	r.LoadHTMLGlob("templates/*")
	//设置静态文件加载目录
	r.Static("/dsb", "./statics")
	r.GET("/login", loginHandler)
	r.GET("/index", indexHandler)

	r.Run()
}
