package repository

import (
	"mqtt_chat_manager/internal/models"

	"gorm.io/gorm"
)

type MessageRepository struct {
	DB *gorm.DB
}

func (r MessageRepository) CreateMessage(message *models.Message) error {
	return r.DB.Create(message).Error
}

func (r MessageRepository) FindByRoomId(roomId int, start_id int, pageSize int) ([]models.Message, error) {
	var messages []models.Message
	err := r.DB.Where("room_id=? AND id < ?", roomId, start_id).Order("id asc").Limit(pageSize).Find(&messages).Error
	return messages, err
}

func (r MessageRepository) DeleteMessage(messageId int) error {
	return r.DB.Where("id=?", messageId).Delete(&models.Message{}).Error
}
