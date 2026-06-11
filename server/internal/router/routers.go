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
		//사용자 생성
		authGroup.POST("/signup", chatHandler.SignUp)
		//사용자 인증
		authGroup.POST("/signin", chatHandler.SignIn)
	}

	// API v1 라우트 그룹
	// 인증 미들웨어 적용
	v1 := r.Group("/api/v1")

	// v1.Use(chatHandler.AuthenticateUser)

	{
		// 대화방 불러오기
		v1.GET("/rooms", chatHandler.GetRooms)
		// 대화방 생성
		v1.POST("/rooms", chatHandler.CreateRoom)
		// 이전 메세지 불러오기
		v1.GET("/rooms/:room_id/messages", chatHandler.GetMessages)
		// 메세지 기록
		v1.POST("/messages", chatHandler.RecordMessage)
	}

	return r
}
