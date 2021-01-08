package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/:name/:id", func(context *gin.Context) {
		name := context.Param("name")
		id := context.Param("id")
		fmt.Printf("id = %v name = %v\n", id, name)
		context.JSON(200, gin.H{"name": name, "id": id})
	})
	engine.Run()
}
