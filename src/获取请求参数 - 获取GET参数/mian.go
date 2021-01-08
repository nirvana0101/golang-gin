package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/user", func(context *gin.Context) {
		username := context.Query("username")
		password := context.DefaultQuery("password", "123456")
		fmt.Printf("username = %v password = %v\n", username, password)
		context.String(http.StatusOK, "%v %v", username, password)
	})
	engine.Run()
}
