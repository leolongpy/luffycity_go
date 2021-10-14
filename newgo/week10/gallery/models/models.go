package models

import "gorm.io/gorm"

type UserInfoModel struct {
	gorm.Model
	Name  string `json:"name"`
	Sex   string `json:"sex"`
	Phone int    `json:"phone"`
	City  string `json:"city"`
}

type TransformedUserInfo struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UserDB struct {
	Db *gorm.DB
}

func NewUsersDB(db *gorm.DB) *UserDB {
	return &UserDB{
		db,
	}
}
