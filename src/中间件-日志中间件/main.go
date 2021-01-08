package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"time"
)

type User struct {
	UserName string    `form:"username" binding:"required"`
	PassWord string    `form:"password" binding:"required"`
	Age      int       `form:"age" binding:"required,gt=10"` //gt 大于10
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	file, _ := os.Create("./gin.log")
	writer := io.MultiWriter(file)
	gin.DefaultWriter = writer
	gin.DefaultErrorWriter = writer
	engine := gin.Default()
	engine.GET("/user", test)
	engine.POST("/user", test)
	engine.Run()
}
func test(context *gin.Context) {
	user := new(User)
	err := context.ShouldBind(user)
	if err != nil {
		log.Println(err)
		context.String(500, fmt.Sprint(err))
		return
	}
	context.String(200, fmt.Sprintf("%v", *user))
}
