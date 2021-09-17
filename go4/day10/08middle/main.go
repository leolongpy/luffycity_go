package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func shopIndexHandler(c *gin.Context) {
	time.Sleep(5 * time.Second)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "shopping/index",
	})
}
func shopHomeHandler(c *gin.Context) {
	time.Sleep(3 * time.Second)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "shopping/home",
	})
}
func castTime(c *gin.Context) {
	start := time.Now()
	c.Set("key", "value")
	c.Next()

	cast := time.Since(start)
	fmt.Println("cast:", cast)
}
func main() {
	//r:=gin.New()
	r := gin.Default() //默认使用两个中间件：1记日志 2.recover

	r.Use(castTime)
	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("/index", shopIndexHandler)
		shoppingGroup.GET("/home", shopHomeHandler)
	}
	r.Run()
}
