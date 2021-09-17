package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func queryStrinigHandler(c *gin.Context) {
	nameVal := c.DefaultQuery("name", "豪杰")
	cityVal := c.Query("city")
	c.JSON(http.StatusOK, gin.H{
		"name": nameVal,
		"city": cityVal,
	})
}

func formHandler(c *gin.Context) {
	nameVal := c.PostForm("name")
	cityVal := c.DefaultPostForm("city", "雄安")
	c.JSON(http.StatusOK, gin.H{
		"name": nameVal,
		"city": cityVal,
	})
}

func paramsHandler(c *gin.Context) {
	actionVal := c.Param("action")
	c.JSON(http.StatusOK, gin.H{
		"action": actionVal,
	})
}
func main() {
	r := gin.Default()
	r.GET("/query_string", queryStrinigHandler)
	r.POST("/form", formHandler)
	r.GET("/book/:action", paramsHandler)
	r.Run()
}
