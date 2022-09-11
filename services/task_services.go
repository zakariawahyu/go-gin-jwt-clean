package services

import (
	"github.com/mashingan/smapping"
	"github.com/zakariawahyu/go-gin-jwt-clean/common/response"
	"github.com/zakariawahyu/go-gin-jwt-clean/dto"
	"github.com/zakariawahyu/go-gin-jwt-clean/entity"
	"github.com/zakariawahyu/go-gin-jwt-clean/repository"
	"github.com/zakariawahyu/go-gin-jwt-clean/validations"
)

type TaskServices interface {
	CreateTask(taskRequest dto.CreateTaskRequest) (*response.TaskResponse, error)
}

type TaskServicesImpl struct {
	taskRepo repository.TaskRepository
}

func NewTaskServices(taskRepository repository.TaskRepository) TaskServices {
	return &TaskServicesImpl{
		taskRepo: taskRepository,
	}
}

func (taskServices *TaskServicesImpl) CreateTask(taskRequest dto.CreateTaskRequest) (*response.TaskResponse, error) {
	var task entity.Task
	if err := validations.ValidateCreateTask(taskRequest); err != nil {
		return nil, err
	}

	if err := smapping.FillStruct(&task, smapping.MapFields(&taskRequest)); err != nil {
		return nil, err
	}

	result, err := taskServices.taskRepo.Create(task)
	if err != nil {
		return nil, err
	}

	res := response.NewTaskResponse(result)
	return &res, nil
}
