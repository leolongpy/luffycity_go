package logic

import (
	"luffycity_go/go4/bloger/db"
	"luffycity_go/go4/bloger/models"
)

func GetArticleRecordList(pageNum, pageSize int) (articleRecordList []*models.ArticleRecord, err error) {
	articleList, err := db.GetArticle(pageNum, pageSize)
	if err != nil {
		return
	}
	if len(articleList) == 0 {
		return
	}
	categoryList, err := db.GetAllCategoryList()
	if err != nil {
		return
	}
	articleRecordList = make([]*models.ArticleRecord, 0, len(articleList))
	for _, a := range articleList {
		for _, c := range categoryList {
			if a.CategoryId == c.CategoryId {
				tmpC := &models.ArticleRecord{
					ArticleInfo: *a,
					Category:    *c,
				}
				articleRecordList = append(articleRecordList, tmpC)
				break
			}
		}
	}
	return
}
