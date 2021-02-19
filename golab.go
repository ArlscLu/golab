package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"姓名": "pong",
		})
	})
	r.GET("/json", func(c *gin.Context) {
		c.AsciiJSON(200, gin.H{
			"姓名": "anmina",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
