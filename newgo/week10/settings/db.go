package settings

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"luffycity_go/newgo/week10/gallery/models"
)

func NewDatabase(cfg *Config) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)
	db, err = gorm.Open(mysql.Open(cfg.DBConfig.DataSourceName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.UserInfoModel{})
	return db, nil
}
