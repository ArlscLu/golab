package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	// 添加 Get 请求路由
	e.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin")
	})
	e.GET("/u", func(context *gin.Context) {
		context.String(http.StatusOK, func() string {
			return "ganbin"
		}())
	})

	e.Run(":8088")
}
