package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zakariawahyu/go-gin-jwt-clean/config"
	"net/http"
)

func main() {
	r := gin.Default()

	config.DatabaseConnection()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})

	r.Run("localhost:8081")
}
