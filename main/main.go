package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//Engin
	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		context.JSON(
			200, gin.H{
				"code":    200,
				"message": "success",
			})
	})

	r.GET("/hello", helloHander)
	r.Run(":8085")
}

func helloHander(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "hello gin",
	})
}
