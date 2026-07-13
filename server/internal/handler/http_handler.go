package handler

import (
	"errors"
	"log"
	"math"
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
	var requestBody models.AuthenticateRequest
	err := ctx.ShouldBindJSON(&requestBody)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "잘못된 요청", "details": err.Error()})
		return
	}

	user, err := h.UserRepo.FindByUserName(requestBody.Username)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(400, gin.H{"error": "사용자명 없음", "details": err.Error()})
			return
		}
		ctx.JSON(500, gin.H{"error": "사용자 조회 실패", "details": err.Error()})
		return
	}

	ctx.JSON(200, user)

}

// 사용자 생성
func (h HttpHandler) SignUp(ctx *gin.Context) {
	var requestBody AuthenticateRequest
	err := ctx.ShouldBindJSON(&requestBody)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "잘못된 요청"})
		return
	}

	// 사용자명 없을시에 생성
	var newUser models.User
	newUser.Username = requestBody.UserName
	err = h.UserRepo.CreateUser(&newUser)

	if err != nil {
		if !errors.Is(err, gorm.ErrDuplicatedKey) {
			ctx.JSON(400, gin.H{"error": "사용자명 중복", "details": err.Error()})
			return
		}

		ctx.JSON(500, gin.H{"error": "사용자 생성 실패",
			"details": err.Error()})
		return
	}
	ctx.JSON(200, newUser)

}

// 메시지 조회
/*
 - room_id: 조회할 대화방 ID
 - start_id: 조회 시작 메시지 ID (이 ID보다 작은 메시지들을 조회)
 - page_size: 한 번에 조회할 메시지 수 (예: 20)
*/
func (h HttpHandler) GetMessages(ctx *gin.Context) {
	roomId := ctx.Param("room_id")
	startId := ctx.Query("start_id")

	roomIdInt, err := strconv.Atoi(roomId)
	startIdInt, err := strconv.Atoi(startId)

	log.Println("roomIdInt: ", roomIdInt, "startIdInt: ", startIdInt)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "파라미터 오류"})
		return
	}

	if startIdInt < 0 {
		startIdInt = math.MaxInt64
	}

	messages, err := h.MessageRepo.FindByRoomId(roomIdInt, startIdInt, 30)

	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	if err != nil {
		ctx.JSON(500, gin.H{"error": "메시지 로드 실패"})
		return
	}
	ctx.JSON(200, messages)
}

// 메시지 기록
func (h HttpHandler) RecordMessage(ctx *gin.Context) {
	var message models.RecordMessageRequest
	err := ctx.ShouldBindJSON(&message)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "잘못된 요청", "details": err.Error()})
		return
	}

	recordMessage := models.Message{
		Id:       message.Id,
		RoomId:   message.RoomId,
		SenderId: message.SenderId,
		Content:  message.Content,
	}

	err = h.MessageRepo.CreateMessage(&recordMessage)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "메시지 저장 실패", "details": err.Error()})
		return
	}
	ctx.JSON(200, recordMessage)
}

// 대화방 생성
func (h HttpHandler) CreateRoom(ctx *gin.Context) {
	var request models.CreateRoomRequest
	err := ctx.ShouldBindJSON(&request)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "잘못된 요청", "details": err.Error()})
		return
	}

	room, err := h.RoomRepo.CreateRoom(request.RoomName)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "방 생성 실패"})
	}

	ctx.JSON(200, room)
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
