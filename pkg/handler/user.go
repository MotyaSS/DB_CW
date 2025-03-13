package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.service.Authorisation.GetAllUsers()
	if err != nil {
		abortWithError(c, err)
		return
	}

	// Возвращаем массив в поле items
	c.JSON(http.StatusOK, gin.H{
		"items": users,
	})
}

func (h *Handler) getUser(ctx *gin.Context) {
	callerId, err := h.getCallerId(ctx)
	if err != nil {
		abortWithError(ctx, err)
		return
	}

	user, err := h.service.Authorisation.GetUserById(callerId)
	if err != nil {
		abortWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *Handler) deleteUser(ctx *gin.Context) {
	callerId, err := h.getCallerId(ctx)
	if err != nil {
		abortWithStatusCode(ctx, http.StatusUnauthorized, "caller id is not provided")
	}

	toDeleteId, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		abortWithStatusCode(ctx, http.StatusBadRequest, "invalid user id")
		return
	}
	if callerId == toDeleteId {
		abortWithStatusCode(ctx, http.StatusBadRequest, "cannot delete own account")
	}

	id := h.service.Authorisation.DeleteUser(toDeleteId)
	ctx.JSON(http.StatusOK, id)
}
