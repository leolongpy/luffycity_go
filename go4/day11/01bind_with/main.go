package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type fromA struct {
	Foo string `json:"foo" xml:"foo" binding:"require"`
}

type fromB struct {
	Bar string `json:"bar" xml:"bar" binding:"require"`
}

type UserInfo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password" binding:"required"`
}

func LoginHander(c *gin.Context) {
	if c.Request.Method == "POST" {
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"username": u.Username,
			"password": u.Password,
		})
	} else {
		c.HTML(http.StatusOK, "login.html", nil)
	}
}

func SomeHander(c *gin.Context) {
	objA := fromA{}
	objB := fromB{}
	// c.ShouldBind 使用了 c.Request.Body,不可重用
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "foo",
		})
	} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "bar",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "nil",
		})
	}
}
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Any("/some", SomeHander)
	r.Any("/login", LoginHander)
	r.Run()

}
