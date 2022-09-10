package services

import (
	"github.com/mashingan/smapping"
	"github.com/zakariawahyu/go-gin-jwt-clean/common/response"
	"github.com/zakariawahyu/go-gin-jwt-clean/dto"
	"github.com/zakariawahyu/go-gin-jwt-clean/entity"
	"github.com/zakariawahyu/go-gin-jwt-clean/repository"
	"github.com/zakariawahyu/go-gin-jwt-clean/validations"
	"log"
)

type UserServices interface {
	FindUserById(id string) (*response.UserResponse, error)
	UpdateUser(request dto.UpdateUserRequest) (*response.UserResponse, error)
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

func (userServices *UserServicesImpl) UpdateUser(request dto.UpdateUserRequest) (*response.UserResponse, error) {
	var user entity.User

	if errValidate := validations.ValidateUpdateUser(request); errValidate != nil {
		return nil, errValidate
	}

	if errMap := smapping.FillStruct(&user, smapping.MapFields(&request)); errMap != nil {
		log.Fatal("failed mapping %v", errMap)
		return nil, errMap
	}

	result, err := userServices.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	res := response.NewUserResponse(result)
	return &res, nil
}
