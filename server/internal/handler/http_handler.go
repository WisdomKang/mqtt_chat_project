package handler

import (
	"mqtt_chat_manager/internal/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
	UserRepo    repository.UserRepository
	MessageRepo repository.MessageRepository
	RoomRepo    repository.RoomRepository
}

func (h HttpHandler) GetMessages(ctx *gin.Context) {
	roomId := ctx.Param("room_id")

	roomIdint, err := strconv.Atoi(roomId)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "파라미터 오류"})
		return
	}

	messages, err := h.MessageRepo.FindByRoomId(int(roomIdint), 0, 20)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "메시지 로드 실패"})
		return
	}
	ctx.JSON(200, messages)
}
