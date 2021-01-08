package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()
	engine.GET("/user/*action", func(context *gin.Context) {
		context.String(200, "hello world")
	})
	engine.Run()
}
