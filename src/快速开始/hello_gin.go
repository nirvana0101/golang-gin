package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()
	engine.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "hello gin",
		})
	})
	engine.Run() //Listening and serving HTTP on :8080
}
