package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	AuthHeader  = "Authorization"
	CallerIdCtx = "caller_id"
)

func (h *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(AuthHeader)

	if header == "" {
		abortWithStatusCode(ctx, http.StatusUnauthorized, "auth header is empty")
		return
	}

	splitHeader := strings.Split(header, " ")
	if len(splitHeader) != 2 {
		abortWithStatusCode(ctx, http.StatusUnauthorized, "incorrect header format")
		return
	}
	userId, err := h.service.ParseToken(splitHeader[1])
	if err != nil {
		abortWithStatusCode(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.Set(CallerIdCtx, userId)
}

func (h *Handler) getCallerId(ctx *gin.Context) (int, error) {
	id, ok := ctx.Get(CallerIdCtx)
	if !ok {
		err := fmt.Errorf("cannot find caller id in context")
		return -1, err
	}
	res, ok := id.(int)
	if !ok {
		err := fmt.Errorf("caller id is not of type int")
		return -1, err
	}
	return res, nil
}
