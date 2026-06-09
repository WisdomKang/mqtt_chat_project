package router

import (
	"mqtt_chat_manager/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(chatHandler *handler.HttpHandler) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")

	{
		v1.GET("/rooms/:room_id/messages", chatHandler.GetMessages)
	}

	return r
}
