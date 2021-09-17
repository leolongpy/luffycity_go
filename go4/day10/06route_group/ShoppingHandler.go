package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func shoppingIndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "shopping/index",
	})
}

func shoppingHomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "shopping/home",
	})
}
