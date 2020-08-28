package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"xorm.io/xorm"
)

func RunXorm() {
	engine, err := xorm.NewEngine("mysql", "devreadonly:ug<H7+hsSmw1@tcp(10.30.60.122:3306)/towngas_coupon_1?charset=utf8")
	if err != nil {
		logrus.Error(err)
	}
	defer engine.Close()
}
