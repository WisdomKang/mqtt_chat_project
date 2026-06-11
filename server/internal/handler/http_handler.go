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

type AuthenticateRequest struct {
	UserName string `json:"username" binding:"required"`
}

// 사용자 인증
func (h HttpHandler) SignIn(ctx *gin.Context) {
	var requestBody AuthenticateRequest
	err := ctx.ShouldBindJSON(&requestBody)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "잘못된 요청", "details": err.Error()})
		return
	}

	user, err := h.UserRepo.FindByUserName(requestBody.UserName)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(400, gin.H{"error": "사용자명 없음", "details": err.Error()})
			return
		}
		ctx.JSON(500, gin.H{"error": "사용자 조회 실패", "details": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"user": user})

}

// 사용자 생성 핸들러
func (h HttpHandler) SignUp(ctx *gin.Context) {
	var request AuthenticateRequest
	err := ctx.ShouldBindJSON(&request)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "잘못된 요청"})
		return
	}

	// 사용자명 없을시에 생성
	var newUser models.User
	newUser.Username = request.UserName
	err = h.UserRepo.CreateUser(newUser)

	if err != nil {
		if !errors.Is(err, gorm.ErrDuplicatedKey) {
			ctx.JSON(400, gin.H{"error": "사용자명 중복", "details": err.Error()})
			return
		}

		ctx.JSON(500, gin.H{"error": "사용자 생성 실패",
			"details": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"user": newUser})

}

// 메시지 조회
func (h HttpHandler) GetMessages(ctx *gin.Context) {
	roomId := ctx.Param("room_id")
	pageIndex := ctx.Param("page_index")

	roomIdint, err := strconv.Atoi(roomId)
	pageIndexint, err := strconv.Atoi(pageIndex)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "파라미터 오류"})
		return
	}

	messages, err := h.MessageRepo.FindByRoomId(int(roomIdint), pageIndexint, 20)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "메시지 로드 실패"})
		return
	}
	ctx.JSON(200, messages)
}

func (h HttpHandler) RecordMessage(ctx *gin.Context) {
	var message models.RecordMessageRequest
	err := ctx.ShouldBindJSON(&message)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "잘못된 요청", "details": err.Error()})
		return
	}

	newMessage := models.Message{
		RoomId:   message.RoomId,
		SenderId: message.SenderId,
		Content:  message.Content,
	}

	err = h.MessageRepo.CreateMessage(newMessage)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "메시지 저장 실패", "details": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": newMessage})
}

// 대화방 생성
func (h HttpHandler) CreateRoom(ctx *gin.Context) {
	roomName := ctx.Param("room_name")
	room, err := h.RoomRepo.CreateRoom(roomName)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "방 생성 실패"})
	}

	ctx.JSON(200, gin.H{"room": room})
}

// 대화방 조회
func (h HttpHandler) GetRooms(ctx *gin.Context) {
	rooms, err := h.RoomRepo.GetRooms()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "방 로드 실패"})
		return
	}
	ctx.JSON(200, rooms)
}
