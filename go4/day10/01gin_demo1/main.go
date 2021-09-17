package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "这是一个index页面",
	})
}
func main() {
	//启动一个默认路由
	rount := gin.Default()
	rount.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "hello",
		})
	})
	rount.GET("/index", indexHandler)
	rount.Run(":9090")
}
