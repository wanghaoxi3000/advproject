package main

import (
	_ "advancedproject/docs"
	"advancedproject/server"
)

// @title API doc
// @version 1.0
// @description advanced project

// @contact.name wanghaoxi3000
// @contact.email wanghaoxi3000@163.com

// @BasePath /api/v1

// https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
func main() {
	r := server.SetupRouter()
	r.Run(":3000")
}
