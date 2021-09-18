package db

import "luffycity_go/go4/bloger/models"

// GetAllCategoryList 查询所有分类信息
func GetAllCategoryList() (categoryList []*models.Category, err error) {
	err = DB.Select(&categoryList, "SELECT id,category_name,category_no FROM category order by category_no asc")
	if err != nil {
		return
	}
	return
}
