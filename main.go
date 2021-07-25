package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "hello gin",
		})
	})
	err := r.Run()
	if err != nil {
		return
	}
}
