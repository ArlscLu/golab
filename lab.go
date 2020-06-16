package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
			"example": [5]uint8{1, 2, 3, 4, 5},
		})
	})
	router.Run(":8080")
}
