package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// cookie示例

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
		if u.Username == "guan" && u.Password == "123" {
			// 登陆成功
			// 1. 设置Cookie
			c.SetCookie("username", u.Username, 20, "/", "127.0.0.1", false, true)
			// 跳转到index页面
			c.Redirect(http.StatusMovedPermanently, toPath)
		} else {
			// 密码错误
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

// 基于Cookie实现用户登录认证的中间件
func cookieMiddleware(c *gin.Context) {
	// 在返回页面之前要先校验是否存在username的Cookie
	// 获取Cookie
	username, err := c.Cookie("username")
	if err != nil {
		// 直接让跳转到登录页面

		toPath := fmt.Sprintf("%s?next=%s", "/login", c.Request.URL.Path)
		c.Redirect(http.StatusMovedPermanently, toPath)
		return
	}
	// 用户已经登录了
	c.Set("username", username) // 在上下文中设置一个键值对
	c.Next()                    // 继续后续的处理函数
}

func homeHandler(c *gin.Context) {

	// 在返回页面之前要先校验是否存在username的Cookie
	// 获取Cookie
	username, err := c.Cookie("username")
	if err != nil {
		// 直接让跳转到登录页面
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
		// 如果取不到值，说明前面中间件出问题了
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
	username, ok := tmpUsername.(string)
	if !ok {
		// 类型断言失败
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
