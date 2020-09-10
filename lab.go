package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

const (
	dsn = "root:123@(127.0.0.1)/test?charset=utf8&parseTime=True&loc=Local"
)

func main() {

}
