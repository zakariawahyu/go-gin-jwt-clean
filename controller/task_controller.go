package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zakariawahyu/go-gin-jwt-clean/common/response"
	"github.com/zakariawahyu/go-gin-jwt-clean/dto"
	"github.com/zakariawahyu/go-gin-jwt-clean/middleware"
	"github.com/zakariawahyu/go-gin-jwt-clean/services"
	"net/http"
	"strconv"
)

type TaskController interface {
	TaskRouters(group *gin.RouterGroup)
	CreateTaskUser(c *gin.Context)
}

type TaskControllerImpl struct {
	taskServices services.TaskServices
	jwtServices  services.JWTServices
}

func NewTaskController(taskServices services.TaskServices, jwtServices services.JWTServices) TaskController {
	return &TaskControllerImpl{
		taskServices: taskServices,
		jwtServices:  jwtServices,
	}
}

func (taskController *TaskControllerImpl) TaskRouters(group *gin.RouterGroup) {
	route := group.Group("/task", middleware.AuthorizeJWT(taskController.jwtServices))
	route.POST("/", taskController.CreateTaskUser)
}

func (taskController *TaskControllerImpl) CreateTaskUser(c *gin.Context) {
	var taskRequest dto.CreateTaskRequest

	if err := c.ShouldBind(&taskRequest); err != nil {
		res := response.BuildErrorResponse("Failed to process request", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	claims := taskController.jwtServices.GetClaimsJWT(c)
	id := fmt.Sprintf("%v", claims["user_id"])
	_id, _ := strconv.ParseInt(id, 0, 64)

	taskRequest.UserID = _id
	result, err := taskController.taskServices.CreateTask(taskRequest)
	if err != nil {
		res := response.BuildErrorResponse("Cant create task", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := response.BuildSuccessResponse("Success", result)
	c.JSON(http.StatusOK, res)
}
