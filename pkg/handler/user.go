package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (*Handler) getAllUsers(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		"all users",
	)
}

func (*Handler) getUser(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		[]string{
			"user ",
			ctx.Param("user_id"),
		},
	)
}

func (*Handler) deleteUser(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		[]string{
			"user deleted",
			ctx.Param("user_id"),
		},
	)
}
