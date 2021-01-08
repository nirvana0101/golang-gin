package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()
	engine.GET("/get", func(context *gin.Context) {
		context.String(200, "Get")
	})
	engine.POST("/post", func(context *gin.Context) {
		context.String(200, "Post")
	})
	engine.DELETE("/delete", func(context *gin.Context) {
		context.String(200, "Delect")
	})
	engine.Any("/any", func(context *gin.Context) {
		context.String(200, "Any")
	})
	engine.Run()
}
