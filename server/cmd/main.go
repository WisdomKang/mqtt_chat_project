package main

import (
	"log"
	"mqtt_chat_manager/internal/database"
	"mqtt_chat_manager/internal/handler"
	"mqtt_chat_manager/internal/repository"
	"mqtt_chat_manager/internal/router"
	"os"
)

var databaseDns = ""

func main() {
	dsn := getEnv("DATABASE_DSN", "host=localhost user=chat_app_user password=test1234 dbname=chat_db port=8094 sslmode=disable TimeZone=Asia/Seoul")
	log.Println(dsn)
	db := database.InitDB(dsn)
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
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
