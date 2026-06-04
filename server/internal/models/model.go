package models

import (
	"time"
)

type User struct {
	UserId   int `gorm:"primaryKey;column:user_id"`
	Username string

	created_at time.Time
	deleted_at *time.Time
}
type Room struct {
	RoomId   int `gorm:"primaryKey;column:room_id"`
	RoomName string

	created_at time.Time
	deleted_at *time.Time
}

type Room_Member struct {
	Room Room `gorm:"foreignKey:room_id"`
	User User `gorm:"foreignKey:user_id"`

	created_at time.Time
	deleted_at *time.Time
}

type Message struct {
	MessageId int `gorm:"primaryKey;column:message_id"`
	RoomId    int
	SenderID  int
	content   string
}
