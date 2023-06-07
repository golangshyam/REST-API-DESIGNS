package main

import (
	_"fmt"
	_"net/http"
	"github.com/gin-gonic/gin"
			)


func main(){

	r:=gin.Default()

	r.GET("/",func(ctx *gin.Context) {
		ctx.String(200,"hello world!")
	})
	r.Run(":8080")
}