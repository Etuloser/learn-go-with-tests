package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hostname": "pong",
		})
	})
	// 指定端口
	r.Run("0.0.0.0:8080")
}
