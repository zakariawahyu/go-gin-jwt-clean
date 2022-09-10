package services

import (
	"errors"
	"fmt"
	"github.com/mashingan/smapping"
	"github.com/zakariawahyu/go-gin-jwt-clean/common/response"
	"github.com/zakariawahyu/go-gin-jwt-clean/dto"
	"github.com/zakariawahyu/go-gin-jwt-clean/repository"
	"github.com/zakariawahyu/go-gin-jwt-clean/validations"
	"gorm.io/gorm"
	"log"
)

type UserServices interface {
	CreateUser(userRequest dto.UserRegisterRequest) (*response.UserResponse, error)
}

type UserServicesImpl struct {
	userRepo repository.UserRepository
}

func NewUserServices(userRepository repository.UserRepository) UserServices {
	return &UserServicesImpl{
		userRepo: userRepository,
	}
}

func (userServices *UserServicesImpl) CreateUser(userRequest dto.UserRegisterRequest) (*response.UserResponse, error) {
	if errValidate := validations.ValidateUser(userRequest); errValidate != nil {
		return nil, errValidate
	}

	user, errFind := userServices.userRepo.FindByEmail(userRequest.Email)
	if errFind == nil {
		return nil, errors.New(fmt.Sprintf("User %v already exists", user.Name))
	}

	if errFind != nil && !errors.Is(errFind, gorm.ErrRecordNotFound) {
		return nil, errFind
	}

	if errMap := smapping.FillStruct(&user, smapping.MapFields(&userRequest)); errMap != nil {
		log.Fatal("failed mapping %v", errMap)
		return nil, errMap
	}

	result, err := userServices.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	res := response.NewUserResponse(result)
	return &res, nil
}
