package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zakariawahyu/go-gin-jwt-clean/common/response"
	"github.com/zakariawahyu/go-gin-jwt-clean/dto"
	"github.com/zakariawahyu/go-gin-jwt-clean/middleware"
	"github.com/zakariawahyu/go-gin-jwt-clean/services"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type TaskController interface {
	TaskRouters(group *gin.RouterGroup)
	CreateTaskUser(c *gin.Context)
	UpdateTaskUser(c *gin.Context)
	GetTaskById(c *gin.Context)
	GetAll(c *gin.Context)
}

type TaskControllerImpl struct {
	taskServices services.TaskServices
	jwtServices  services.JWTServices
	userServices services.UserServices
}

func NewTaskController(taskServices services.TaskServices, jwtServices services.JWTServices, userServices services.UserServices) TaskController {
	return &TaskControllerImpl{
		taskServices: taskServices,
		jwtServices:  jwtServices,
		userServices: userServices,
	}
}

func (taskController *TaskControllerImpl) TaskRouters(group *gin.RouterGroup) {
	route := group.Group("/task", middleware.AuthorizeJWT(taskController.jwtServices))
	route.POST("/", taskController.CreateTaskUser)
	route.PUT("/:id", taskController.UpdateTaskUser)
	route.GET("/:id", taskController.GetTaskById)
	route.GET("/", taskController.GetAll)
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

func (taskController *TaskControllerImpl) UpdateTaskUser(c *gin.Context) {
	var taskRequest dto.UpdateTaskRequest

	if err := c.ShouldBind(&taskRequest); err != nil {
		res := response.BuildErrorResponse("Failed to process request", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	claims := taskController.jwtServices.GetClaimsJWT(c)
	userId := fmt.Sprintf("%v", claims["user_id"])
	_userId, _ := strconv.ParseInt(userId, 0, 64)

	taskId := c.Param("id")
	_taskId, _ := strconv.ParseInt(taskId, 0, 64)

	taskRequest.ID = _taskId
	taskRequest.UserID = _userId
	result, err := taskController.taskServices.UpdateTask(taskRequest)
	if err != nil {
		res := response.BuildErrorResponse("Cant update task", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := response.BuildSuccessResponse("Success", result)
	c.JSON(http.StatusOK, res)
}

func (taskController *TaskControllerImpl) GetTaskById(c *gin.Context) {
	claims := taskController.jwtServices.GetClaimsJWT(c)
	userId := fmt.Sprintf("%v", claims["user_id"])
	taskId := c.Param("id")

	result, err := taskController.taskServices.FindTaskById(taskId, userId)
	if err != nil {
		res := response.BuildErrorResponse("Cant get task", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, res)
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		}
		return
	}

	res := response.BuildSuccessResponse("Success", result)
	c.JSON(http.StatusOK, res)
}

func (taskController *TaskControllerImpl) GetAll(c *gin.Context) {
	claims := taskController.jwtServices.GetClaimsJWT(c)
	userId := fmt.Sprintf("%v", claims["user_id"])

	result, err := taskController.taskServices.GetAllTask(userId)
	if err != nil {
		res := response.BuildErrorResponse("Cant update task", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := response.BuildSuccessResponse("Success", result)
	c.JSON(http.StatusOK, res)
}
