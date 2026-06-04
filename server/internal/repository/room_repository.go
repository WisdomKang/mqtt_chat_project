package repository

import (
	"mqtt_chat_manager/internal/models"

	"gorm.io/gorm"
)

type RoomRepository struct {
	DB *gorm.DB
}

func (r RoomRepository) CreateRoom(roomName string) error {
	return r.DB.Create(&models.Room{RoomName: roomName}).Error
}

func (r RoomRepository) FindByRoomName(roomName string) (*models.Room, error) {
	var room models.Room
	err := r.DB.Where("room_name=?", roomName).First(&room).Error
	return &room, err
}
func (r RoomRepository) FindByRoomId(roomId int, page int) (*models.Room, error) {
	var room models.Room
	err := r.DB.Where("room_id=?", roomId).First(&room).Error
	return &room, err
}
