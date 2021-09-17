package main

import "github.com/gin-gonic/gin"

// 路由分组
func main() {
	r := gin.Default()
	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("/index", shoppingIndexHandler)
		shoppingGroup.GET("/home", shoppingHomeHandler)
	}

	liveGroup := r.Group("/live")
	{
		liveGroup.GET("/index", liveIndexHandler)
		liveGroup.GET("/home", liveHomeHandler)
	}
	r.Run()
}
