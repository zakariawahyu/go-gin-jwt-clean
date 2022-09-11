package repository

import (
	"github.com/zakariawahyu/go-gin-jwt-clean/entity"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task entity.Task) (entity.Task, error)
}

type TaskRepositoryImpl struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &TaskRepositoryImpl{
		db: db,
	}
}

func (taskRepo *TaskRepositoryImpl) Create(task entity.Task) (entity.Task, error) {
	if err := taskRepo.db.Create(&task).Error; err != nil {
		return task, err
	}

	taskRepo.db.Preload("User").Find(&task)
	return task, nil
}
