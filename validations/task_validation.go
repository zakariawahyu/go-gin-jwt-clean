package validations

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/zakariawahyu/go-gin-jwt-clean/dto"
)

func ValidateCreateTask(task dto.CreateTaskRequest) error {
	return validation.ValidateStruct(&task,
		validation.Field(&task.Title, validation.Required, validation.Length(10, 0)),
		validation.Field(&task.Description, validation.Required),
		validation.Field(&task.UserID, validation.Required))
}
