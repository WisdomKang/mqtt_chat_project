package models

import (
	"time"
)

type User struct {
	UserId   int    `gorm:"primaryKey;column:user_id" json:"user_id"`
	Username string `json:"username"`

	created_at time.Time
	deleted_at *time.Time
}
type Room struct {
	RoomId   int    `gorm:"primaryKey;column:room_id" json:"room_id"`
	RoomName string `json:"room_name"`

	// RoomMembers []User `gorm:"many2many:room_members;foreignKey:RoomId;joinForeignKey:room_id;References:UserId;joinReferences:user_id" json:"room_members"`

	created_at time.Time
	deleted_at *time.Time
}

// type RoomMember struct {
// 	MemberId int  `gorm:"primaryKey;column:member_id"`
// 	Room     Room `gorm:"foreignKey:room_id"`
// 	User     User `gorm:"foreignKey:user_id"`

// 	created_at time.Time
// 	deleted_at *time.Time
// }

type Message struct {
	MessageId int64  `gorm:"primaryKey;column:message_id" json:"message_id"`
	RoomId    int    `gorm:"column:room_id" json:"room_id"`
	SenderId  int    `gorm:"column:sender_id" json:"sender_id"`
	Content   string `gorm:"column:content" json:"content"`

	created_at time.Time
}
