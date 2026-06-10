package models

import (
	"time"
)

type User struct {
	UserId   int    `gorm:"primaryKey;column:user_id" ,json:"user_id"`
	Username string `json:"username"`

	created_at time.Time
	deleted_at *time.Time
}
type Room struct {
	RoomId   int    `gorm:"primaryKey;column:room_id" json:"room_id"`
	RoomName string `json:"room_name"`

	RommMembers []User `gorm:"many2many:room_members;foreignKey:RoomId;joinForeignKey:room_id;References:UserId;joinReferences:user_id"`

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
	MessageId int    `gorm:"primaryKey;column:message_id"`
	RoomId    int    `gorm:"column:room_id"`
	SenderID  int    `gorm:"column:sender_id"`
	content   string `gorm:"column:content"`
}
