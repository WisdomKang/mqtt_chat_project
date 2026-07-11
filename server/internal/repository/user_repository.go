package repository

import (
	"mqtt_chat_manager/internal/models"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r UserRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r UserRepository) FindByUserName(username string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("username=?", username).First(&user).Error

	return &user, err
}

func (r UserRepository) DeleteUser(userId int) error {
	return r.DB.Where("id=?", userId).Update("deleted_at", time.Now()).Error
}
