package handler

import (
	"math"
	"net/http"

	entity "github.com/MotyaSS/DB_CW/pkg/entities"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllRoles(ctx *gin.Context) {
	res, err := h.service.GetAllRoles()
	if err != nil {
		abortWithError(ctx, err)
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) signUp(ctx *gin.Context) {
	var input entity.User

	if err := ctx.BindJSON(&input); err != nil {
		abortWithStatusCode(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Authorisation.CreateCustomer(input)
	if err != nil {
		abortWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]any{
		"id": id,
	})
}

func (h *Handler) signUpPrivileged(ctx *gin.Context) {
	callerId, err := h.getCallerId(ctx)
	if err != nil {
		abortWithStatusCode(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var input entity.User
	input.RoleId = math.MinInt
	if err := ctx.BindJSON(&input); err != nil {
		abortWithStatusCode(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if input.RoleId == math.MinInt {
		abortWithStatusCode(ctx, http.StatusBadRequest, "no role provided")
		return
	}

	id, err := h.service.Authorisation.CreateUser(callerId, input)
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
		abortWithError(ctx, err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		map[string]any{
			"token": token,
		})
}
