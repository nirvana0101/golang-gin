package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func main() {
	engine := gin.Default()
	engine.POST("/user", func(context *gin.Context) {
		bytes, e := ioutil.ReadAll(context.Request.Body)
		if e != nil {
			fmt.Println("err = ", e)
		}
		context.String(200, string(bytes))
	})
	engine.Run()
}
