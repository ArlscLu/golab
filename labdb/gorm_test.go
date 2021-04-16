package labdb

import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gormDb *gorm.DB

func init() {
	dsn := "test_minqijia:test_minqijia_rw@tcp(10.20.1.20:3306)/towngas_community?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	gormDb = db
}

func TestConn(t *testing.T) {
	first := TaskConfig{
		GroupID: 1,
	}
	_ = gormDb.First(&first)
	t.Log(first)
}
