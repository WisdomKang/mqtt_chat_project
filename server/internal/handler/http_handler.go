package handler

import (
	"errors"
	"mqtt_chat_manager/internal/models"
	"mqtt_chat_manager/internal/repository"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

// 사용자 생성 핸들러
func (h HttpHandler) CreateUserId(ctx *gin.Context) {
	userName := ctx.Param("user_name")

	// 사용자 이름으로 조회
	_, err := h.UserRepo.FindByUserName(userName)

	// 사용자명 이미 있을시 중복 생성 방지
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(500, gin.H{"error": err})
			return
		}
	}

	// 사용자명 없을시에 생성
	var newUser models.User
	err = h.UserRepo.CreateUser(newUser)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "사용자 생성 실패"})
	}
	ctx.JSON(200, gin.H{"user": newUser})

}

func (h HttpHandler) CreateRoom(ctx *gin.Context) {
	roomName := ctx.Param("room_name")
	room, err := h.RoomRepo.CreateRoom(roomName)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "방 생성 실패"})
	}

	ctx.JSON(200, gin.H{"room": room})
}

func (h HttpHandler) GetRooms(ctx *gin.Context) {
	rooms, err := h.RoomRepo.GetRooms()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "방 로드 실패"})
		return
	}
	ctx.JSON(200, rooms)
}
