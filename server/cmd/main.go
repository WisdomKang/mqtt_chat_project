package main

import (
	"mqtt_chat_manager/internal/database"
	"mqtt_chat_manager/internal/models"
	"mqtt_chat_manager/internal/repository"
)

func main() {
	db := database.InitDB()

	repo := repository.UserRepository{DB: db}

	repo.CreateUser(models.User{Username: "강현명!"})
	// repo.DeleteUser(2)
}
