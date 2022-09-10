package repository

import (
	"github.com/zakariawahyu/go-gin-jwt-clean/entity"
	"github.com/zakariawahyu/go-gin-jwt-clean/exception"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user entity.User) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	FindById(id string) (entity.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: database,
	}
}

func (repository *UserRepositoryImpl) Create(user entity.User) (entity.User, error) {
	user.Password = hashAndSalt([]byte(user.Password))
	if err := repository.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	if err := repository.db.Where("email = ?", email).Take(&user); err != nil {
		return user, err.Error
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindById(id string) (entity.User, error) {
	var user entity.User
	if err := repository.db.Where("id = ?", id).Take(&user); err != nil {
		return user, err.Error
	}

	return user, nil
}

func hashAndSalt(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	exception.PanicIfNeeded(err)
	return string(hash)
}
