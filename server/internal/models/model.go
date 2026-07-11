package models

import (
	"time"
)

type User struct {
	Id       int    `gorm:"primaryKey;column:id" json:"id"`
	Username string `json:"username"`

	created_at time.Time
	deleted_at *time.Time
}
type Room struct {
	Id       int    `gorm:"primaryKey;column:id" json:"id"`
	RoomName string `json:"room_name"`

	created_at time.Time
	deleted_at *time.Time
}

type Message struct {
	Id       int64  `gorm:"primaryKey;column:id" json:"id"`
	RoomId   int    `gorm:"column:room_id" json:"room_id"`
	SenderId int    `gorm:"column:sender_id" json:"sender_id"`
	Content  string `gorm:"column:content" json:"content"`

	created_at time.Time
}
