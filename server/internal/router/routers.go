package router

import (
	"mqtt_chat_manager/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(chatHandler *handler.HttpHandler) *gin.Engine {
	r := gin.New()

	// 인증 관련 라우트 그룹
	authGroup := r.Group("/auth")

	{
		authGroup.POST("/signup", chatHandler.CreateUserId)
		authGroup.POST("/signin", chatHandler.AuthenticateUser)
	}

	// API v1 라우트 그룹
	// 인증 미들웨어 적용
	v1 := r.Group("/api/v1")

	v1.Use(chatHandler.AuthenticateUser)

	{
		// 대화방 불러오기
		v1.GET("/rooms", chatHandler.GetRooms)

		// 이전 메세지 불러오기
		v1.GET("/rooms/:room_id/messages", chatHandler.GetMessages)
	}

	return r
}
