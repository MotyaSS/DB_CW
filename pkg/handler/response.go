package handler

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func abortWithStatusCode(ctx *gin.Context, statusCode int, msg string) {
	slog.Error(msg)
	ctx.AbortWithStatusJSON(statusCode, map[string]string{"message": msg})
}
