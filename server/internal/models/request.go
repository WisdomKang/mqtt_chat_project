package models

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
}

type CreateRoomRequest struct {
	RoomName string `json:"room_name" binding:"required"`
}

type GetRoomsRequest struct {
	UserId int `json:"user_id" binding:"required"`
}

type GetMessagesRequest struct {
	RoomId int `json:"room_id" binding:"required"`
}

type RecordMessageRequest struct {
	RoomId   int    `json:"room_id" binding:"required"`
	SenderId int    `json:"user_id" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

type AuthenticateRequest struct {
	Username string `json:"username" binding:"required"`
}
