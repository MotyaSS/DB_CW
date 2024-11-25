package handler

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type responseError struct {
	Message string `json:"message"`
}

func abortWithStatusCode(ctx *gin.Context, statusCode int, msg string) {
	slog.Error(msg)
	// alternative is to use map[string]string but ig it takes more time and space than simple struct
	// I might be wrong abt that one 0_()
	ctx.AbortWithStatusJSON(statusCode, responseError{msg})
}
