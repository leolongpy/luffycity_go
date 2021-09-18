package logic

import (
	"luffycity_go/go4/bloger/db"
	"luffycity_go/go4/bloger/models"
)

func GetCategoryList() (categoryList []*models.Category, err error) {
	return db.GetAllCategoryList()
}
