package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("template/*")
	engine.GET("/", func(context *gin.Context) {
		context.HTML(200, "index.html", gin.H{
			"title": "这是index页面",
		})

	})
	engine.Run()

}
