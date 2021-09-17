package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func bookListHanadler(c *gin.Context) {
	bookList, err := QueryAllBook()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}
	c.HTML(http.StatusOK, "book/book_list.tmpl", gin.H{
		"code": 0,
		"data": bookList,
	})
}
func newBookHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "book/new_book.html", nil)
}

func createBookHandler(c *gin.Context) {
	var msg string
	titleVal := c.PostForm("title")
	priceVal := c.PostForm("price")
	price, err := strconv.ParseFloat(priceVal, 64)
	if err != nil {
		msg = "无效的价格参数"
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
		return
	}
	fmt.Println("%T %T\n", titleVal, priceVal)
	err = InsertBook(titleVal, price)
	if err != nil {
		msg = "插入数据失败，请重试！"
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/book/list")
}

func deleteBookHandler(c *gin.Context) {
	idStr := c.Query("id")
	idVal, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}
	err = DeleteBook(idVal)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
	}
	c.Redirect(http.StatusMovedPermanently, "/book/list")
}

func main() {
	err := InitDB()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/book/list", bookListHanadler)
	r.GET("/book/new", newBookHandler)
	r.POST("/book/new", createBookHandler)
	r.GET("/book/delete", deleteBookHandler)
	r.Run()
}
