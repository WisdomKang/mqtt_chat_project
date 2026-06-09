package main

import (
	"mqtt_chat_manager/internal/database"
	"mqtt_chat_manager/internal/handler"
	"mqtt_chat_manager/internal/repository"
	"mqtt_chat_manager/internal/router"
)

func main() {
	db := database.InitDB()
	messageRepo := repository.MessageRepository{DB: db}
	userRepo := repository.UserRepository{DB: db}
	roomRepo := repository.RoomRepository{DB: db}

	chatHandler := handler.HttpHandler{
		UserRepo:    userRepo,
		MessageRepo: messageRepo,
		RoomRepo:    roomRepo,
	}

	router := router.SetupRouter(&chatHandler)

	router.Run(":8080")

}
