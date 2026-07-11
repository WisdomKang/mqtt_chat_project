package models

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
}

type CreateRoomRequest struct {
	RoomName string `json:"room_name" binding:"required"`
}

type GetRoomsRequest struct {
	Id int `json:"id" binding:"required"`
}

type JoinRoomRequest struct {
	Id     int `json:"id" binding:"required"`
	RoomId int `json:"room_id" binding:"required"`
}

type GetMessagesRequest struct {
	Id int `json:"id" binding:"required"`
}

type RecordMessageRequest struct {
	Id       int64  `json:"id" binding:"required"`
	RoomId   int    `json:"room_id" binding:"required"`
	SenderId int    `json:"sender_id" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

type AuthenticateRequest struct {
	Username string `json:"username" binding:"required"`
}
