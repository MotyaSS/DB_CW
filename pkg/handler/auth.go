package handler

import (
	"net/http"

	entity "github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var input entity.User

	if err := ctx.BindJSON(&input); err != nil {
		abortWithStatusCode(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// TODO: change logic here to handle with CreateUser method and handle request with no role_id
	id, err := h.service.Authorisation.CreateCustomer(input)
	if err != nil {
		abortWithStatusCode(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]any{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(ctx *gin.Context) {
	var input signInInput
	if err := ctx.Bind(&input); err != nil {
		abortWithStatusCode(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Authorisation.GenerateToken(input.Username, input.Password)
	if err != nil {
		abortWithStatusCode(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(
		http.StatusOK,
		map[string]any{
			"token": token,
		})
}
