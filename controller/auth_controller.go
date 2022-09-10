package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zakariawahyu/go-gin-jwt-clean/common/response"
	"github.com/zakariawahyu/go-gin-jwt-clean/dto"
	"github.com/zakariawahyu/go-gin-jwt-clean/services"
	"net/http"
)

type AuthController interface {
	AuthRoutes(group *gin.RouterGroup)
	Register(c *gin.Context)
}

type AuthControllerImpl struct {
	userServices services.UserServices
}

func NewAuthController(userServices services.UserServices) AuthController {
	return &AuthControllerImpl{
		userServices: userServices,
	}
}

func (authController *AuthControllerImpl) AuthRoutes(group *gin.RouterGroup) {
	route := group.Group("/auth")
	route.POST("/register", authController.Register)
}

func (authController *AuthControllerImpl) Register(c *gin.Context) {
	var registerRequest dto.UserRegisterRequest

	if err := c.ShouldBind(&registerRequest); err != nil {
		res := response.BuildErrorResponse("Failed to process request", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := authController.userServices.CreateUser(registerRequest)
	if err != nil {
		res := response.BuildErrorResponse("Cant create user", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := response.BuildResponse(true, "Success", result)
	c.JSON(http.StatusCreated, res)
}
