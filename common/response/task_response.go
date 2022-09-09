package response

import "github.com/zakariawahyu/go-gin-jwt-clean/entity"

type TaskResponse struct {
	ID          int64        `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	User        UserResponse `json:"user"`
}

func NewTaskResponse(task entity.Task) TaskResponse {
	return TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		User:        NewUserResponse(task.User),
	}
}

func NewTaskResponseArray(task []entity.Task) []TaskResponse {
	taskRes := []TaskResponse{}
	for _, value := range task {
		task := TaskResponse{
			ID:          value.ID,
			Title:       value.Title,
			Description: value.Description,
			User:        NewUserResponse(value.User),
		}
		taskRes = append(taskRes, task)
	}
	return taskRes
}
