package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zakariawahyu/go-gin-jwt-clean/config"
	"github.com/zakariawahyu/go-gin-jwt-clean/controller"
	"github.com/zakariawahyu/go-gin-jwt-clean/repository"
	"github.com/zakariawahyu/go-gin-jwt-clean/services"
	"net/http"
)

func main() {
	db := config.DatabaseConnection()
	userRepository := repository.NewUserRepository(db)
	_ = services.NewUserServices(userRepository)
	jwtServices := services.NewJWTServices()
	authServices := services.NewAuthServices(userRepository)
	authController := controller.NewAuthController(authServices, jwtServices)

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})

	// Setup Routers
	v1 := r.Group("api/v1")
	authController.AuthRoutes(v1)

	r.Run("localhost:8081")
}
