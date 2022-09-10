package validations

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/zakariawahyu/go-gin-jwt-clean/dto"
)

func ValidateUpdateUser(user dto.UpdateUserRequest) error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Name, validation.Required),
		validation.Field(&user.Email, validation.Required),
		validation.Field(&user.Password, validation.Length(6, 0)))
}
