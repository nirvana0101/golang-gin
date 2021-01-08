package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()
	engine.Static("/static", "./static") /** ./表示当前文件夹 */
	engine.Run()
}
