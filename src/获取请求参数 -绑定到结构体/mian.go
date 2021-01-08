package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type User struct {
	UserName string    `form:"username"`
	PassWord string    `form:"password"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	engine := gin.Default()
	engine.GET("/user", test)
	engine.POST("/user", test)
	engine.Run()
}
func test(context *gin.Context) {
	user := new(User)
	err := context.ShouldBind(user)
	if err == nil {
		fmt.Println(user.UserName, user.PassWord)
	}
	context.String(200, fmt.Sprintf("%v", *user))
}
