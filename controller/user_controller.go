package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/zakariawahyu/go-gin-jwt-clean/common/response"
	"github.com/zakariawahyu/go-gin-jwt-clean/dto"
	"github.com/zakariawahyu/go-gin-jwt-clean/middleware"
	"github.com/zakariawahyu/go-gin-jwt-clean/services"
	"net/http"
	"strconv"
)

type UserController interface {
	UserRoutes(group *gin.RouterGroup)
	Profile(c *gin.Context)
	Update(c *gin.Context)
}

type UserControllerImpl struct {
	userServices services.UserServices
	jwtServices  services.JWTServices
}

func NewUserController(userServices services.UserServices, jwtServices services.JWTServices) UserController {
	return &UserControllerImpl{
		userServices: userServices,
		jwtServices:  jwtServices,
	}
}

func (userController *UserControllerImpl) UserRoutes(group *gin.RouterGroup) {
	router := group.Group("/user", middleware.AuthorizeJWT(userController.jwtServices))
	router.GET("/profile", userController.Profile)
	router.PUT("/profile", userController.Update)
}

func (userController *UserControllerImpl) Profile(c *gin.Context) {
	header := c.GetHeader("Authorization")
	token := userController.jwtServices.ValidateToken(header, c)
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])

	user, err := userController.userServices.FindUserById(id)
	if err != nil {
		res := response.BuildErrorResponse("Error", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
	}

	res := response.BuildSuccessResponse("Success", user)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func (userController *UserControllerImpl) Update(c *gin.Context) {
	var userRequest dto.UpdateUserRequest

	if err := c.ShouldBind(&userRequest); err != nil {
		res := response.BuildErrorResponse("Failed to process request", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	header := c.GetHeader("Authorization")
	token := userController.jwtServices.ValidateToken(header, c)
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])

	_id, _ := strconv.ParseInt(id, 0, 64)
	userRequest.ID = _id
	result, err := userController.userServices.UpdateUser(userRequest)
	if err != nil {
		res := response.BuildErrorResponse("Error", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := response.BuildSuccessResponse("Success", result)
	c.JSON(http.StatusOK, res)
}
