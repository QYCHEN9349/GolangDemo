package main

import "github.com/gin-gonic/gin"

type Response struct {
	Code int
	Msg  string
	Data any
}

func Index(ctx *gin.Context) {
	ctx.JSONP(200, Response{
		200,
		"success",
		1,
	})
}
func main() {
	//init
	engine := gin.Default()
	// route
	engine.GET("/index", Index)
	engine.Run(":8080")
}
