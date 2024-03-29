package main

import (
	"github.com/gin-gonic/gin"
	ginsession "luffycity_go/go4/day12/gin-session"
	"net/http"
)

// 测试Session服务的gin demo
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// session中间件应该作为一个全局的中间件
	// 初始化全局的MgrObj对象
	ginsession.InitMgr("redis", "127.0.0.1:6379") // 后面可以扩展Redis\mango版
	option := &ginsession.Option{
		MaxAge:   600,
		Path:     "/",
		Domain:   "127.0.0.1",
		Secure:   false,
		HttpOnly: true,
	}
	r.Use(ginsession.SessionMiddleware(ginsession.MgrObj, option))
	r.Any("/login", loginHandler)
	r.GET("/index", indexHandler)
	r.GET("/home", homeHandler)
	r.GET("/vip", AuthMiddleware, vipHandler)
	// 没有匹配的路由都走这个
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	r.Run()

}
