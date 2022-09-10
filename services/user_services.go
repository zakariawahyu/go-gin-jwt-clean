package services

import (
	"github.com/zakariawahyu/go-gin-jwt-clean/repository"
)

type UserServices interface {
}

type UserServicesImpl struct {
	userRepo repository.UserRepository
}

func NewUserServices(userRepository repository.UserRepository) UserServices {
	return &UserServicesImpl{
		userRepo: userRepository,
	}
}
