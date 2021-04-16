package tgorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	// dsn := "test_minqijia:test_minqijia_rw@tcp(10.20.1.20:3306)/towngas_community?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "go:!tewe&213QQ@tcp(127.0.0.1:3306)/blog"
	gormDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = gormDb
}
