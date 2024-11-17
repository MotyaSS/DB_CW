package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func signUp(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, "user signed up")
}

func signIn(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "user signed in")
}
