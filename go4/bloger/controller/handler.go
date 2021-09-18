package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"luffycity_go/go4/bloger/logic"
	"net/http"
	"strconv"
)

func IndexHandler(c *gin.Context) {
	var (
		pageNum  int
		pageSize int
	)
	tmpPageNum, _ := strconv.ParseInt(c.DefaultQuery("page_num", "0"), 10, 64)
	pageNum = int(tmpPageNum)

	tmpPageSize, _ := strconv.ParseInt(c.DefaultQuery("page_size", "0"), 10, 64)
	pageSize = int(tmpPageSize)
	articleRecordList, err := logic.GetArticleRecordList(pageNum, pageSize)
	if err != nil {
		fmt.Printf("get articlerecord List failed,err:%v\n", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
	}
	categoryList, err := logic.GetCategoryList()
	if err != nil {
		fmt.Printf("get category list failed, err:%v\n", err)
	}
	var data = make(map[string]interface{}, 6)
	data["article_list"] = articleRecordList
	data["category_list"] = categoryList
	c.HTML(http.StatusOK, "views/index.html", data)
}
