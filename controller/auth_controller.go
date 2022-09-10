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
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type AuthControllerImpl struct {
	authServices services.AuthServices
}

func NewAuthController(authServices services.AuthServices) AuthController {
	return &AuthControllerImpl{
		authServices: authServices,
	}
}

func (authController *AuthControllerImpl) AuthRoutes(group *gin.RouterGroup) {
	route := group.Group("/auth")
	route.POST("/login", authController.Login)
	route.POST("/register", authController.Register)
}

func (authController *AuthControllerImpl) Register(c *gin.Context) {
	var registerRequest dto.RegisterRequest

	if err := c.ShouldBind(&registerRequest); err != nil {
		res := response.BuildErrorResponse("Failed to process request", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := authController.authServices.RegisterUser(registerRequest)
	if err != nil {
		res := response.BuildErrorResponse("Cant create user", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := response.BuildSuccessResponse("Success", result)
	c.JSON(http.StatusCreated, res)
}

func (authController *AuthControllerImpl) Login(c *gin.Context) {
	var loginRequest dto.LoginRequest

	if err := c.ShouldBind(&loginRequest); err != nil {
		res := response.BuildErrorResponse("Failed to process request", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	user, err := authController.authServices.VerifyCredential(loginRequest)
	if err != nil {
		res := response.BuildErrorResponse("Failed to login", err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	res := response.BuildSuccessResponse("Success", user)
	c.JSON(http.StatusOK, res)
}
