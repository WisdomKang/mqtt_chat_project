package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		log.Println("[🔐 Middleware] Authorization 헤더:", token)
	}
}
