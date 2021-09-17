package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func loginHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		toPath := c.DefaultQuery("next", "/index")
		var u UserInfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"err": "用户名或密码不能为空",
			})
			return
		}
		if u.Username == "leo" && u.Password == "123" {
			c.SetCookie("username", u.Username, 20, "/", "127.0.0.1", false, true)
			c.Redirect(http.StatusMovedPermanently, toPath)
		} else {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"err": "用户名或密码错误",
			})
			return
		}
	} else {
		c.HTML(http.StatusOK, "login.html", nil)
	}
}
func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

//登录中间件
func cookieMiddleware(c *gin.Context) {
	username, err := c.Cookie("username")
	if err != nil {
		toPath := fmt.Sprintf("%s?next=%s", "/login", c.Request.URL.Path)
		c.Redirect(http.StatusMovedPermanently, toPath)
		return
	}
	c.Set("username", username)
	c.Next()
}
func homeHandler(c *gin.Context) {
	username, err := c.Cookie("username")
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
	c.HTML(http.StatusOK, "home.html", gin.H{
		"username": username,
	})
}
func vipHandler(c *gin.Context) {
	tmpUsername, ok := c.Get("username")
	if !ok {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
	username, ok := tmpUsername.(string)
	if !ok {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
	c.HTML(http.StatusOK, "vip.html", gin.H{
		"username": username,
	})
}
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Any("/login", loginHandler)
	r.GET("/index", indexHandler)
	r.GET("/home", homeHandler)
	r.GET("/vip", cookieMiddleware, vipHandler)
	r.Run()
}
