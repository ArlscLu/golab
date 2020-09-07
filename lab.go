package main

import (
	"os"

	"github.com/arlsclu7/golab/std"
	"github.com/gin-gonic/gin"
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
	os.Exit(3)
	std.Main()
	os.Exit(2)
	std.RunLog()
	os.Exit(1)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/post", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"code":    1,
			"data":    "",
		})
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080

}
