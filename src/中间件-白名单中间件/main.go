package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	engine.Use(whitelist)
	engine.GET("/hello", func(context *gin.Context) {
		context.String(200, "hello world")
	})
	engine.Run()
}

func whitelist(context *gin.Context) {
	whiteList := []string{"127.0.0.2", "127.0.0.3"}
	clientIp := context.ClientIP()
	passFlag := false
	for _, ip := range whiteList {
		if ip == clientIp {
			passFlag = true
			break
		}
	}
	if !passFlag {
		context.String(200, fmt.Sprintf("%v not in whitelist", clientIp))
		context.Abort()
	}
}
