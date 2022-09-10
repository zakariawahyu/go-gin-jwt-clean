package services

import (
	"github.com/zakariawahyu/go-gin-jwt-clean/common/response"
	"github.com/zakariawahyu/go-gin-jwt-clean/repository"
)

type UserServices interface {
	FindUserById(id string) (*response.UserResponse, error)
}

type UserServicesImpl struct {
	userRepo repository.UserRepository
}

func NewUserServices(userRepository repository.UserRepository) UserServices {
	return &UserServicesImpl{
		userRepo: userRepository,
	}
}

func (userServices *UserServicesImpl) FindUserById(id string) (*response.UserResponse, error) {
	user, err := userServices.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	result := response.NewUserResponse(user)
	return &result, nil
}
