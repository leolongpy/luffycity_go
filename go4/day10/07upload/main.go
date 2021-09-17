package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func uploadHandler(c *gin.Context) {
	fileObj, err := c.FormFile("filename")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
		return
	}
	filePath := fmt.Sprintf("./%s", fileObj.Filename)
	c.SaveUploadedFile(fileObj, filePath)
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./upload.html")
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", nil)
	})
	r.POST("/upload", uploadHandler)
	r.Run()
}
