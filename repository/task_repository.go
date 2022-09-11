package repository

import (
	"errors"
	"github.com/zakariawahyu/go-gin-jwt-clean/entity"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task entity.Task) (entity.Task, error)
	Update(task entity.Task) (entity.Task, error)
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

func (taskRepo *TaskRepositoryImpl) Update(task entity.Task) (entity.Task, error) {
	result := taskRepo.db.Where("id = ? AND user_id = ?", task.ID, task.UserID).Updates(&task)

	if result.RowsAffected == 0 {
		return task, errors.New("You dont have access to update this task")
	}

	taskRepo.db.Preload("User").Find(&task)
	return task, nil
}
