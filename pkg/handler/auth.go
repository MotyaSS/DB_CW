package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (*Handler) signUp(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, "user signed up")
}

func (*Handler) signIn(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "user signed in")
}
