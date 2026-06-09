package repository

import (
	"mqtt_chat_manager/internal/models"

	"gorm.io/gorm"
)

type RoomRepository struct {
	DB *gorm.DB
}

func (r RoomRepository) CreateRoom(roomName string) (*models.Room, error) {
	var room = models.Room{
		RoomName: roomName,
	}
	err := r.DB.Create(&room).Error
	return &room, err
}

func (r RoomRepository) FindByRoomName(roomName string) (*models.Room, error) {
	var room models.Room
	err := r.DB.Where("room_name=?", roomName).First(&room).Error
	return &room, err
}

func (r RoomRepository) GetRooms() ([]models.Room, error) {
	var rooms []models.Room
	err := r.DB.Find(&rooms).Error
	return rooms, err
}

func (r RoomRepository) FindByRoomId(roomId int, page int) (*models.Room, error) {
	var room models.Room
	err := r.DB.Where("room_id=?", roomId).First(&room).Error
	return &room, err
}
