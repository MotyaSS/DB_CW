package handler

import (
	"net/http"
	"strings"

	"github.com/MotyaSS/DB_CW/pkg/httpError"
	"github.com/gin-gonic/gin"
)

const (
	AuthHeader  = "Authorization"
	CallerIdCtx = "caller_id"
)

func (h *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(AuthHeader)
	if header == "" {
		abortWithStatusCode(ctx, http.StatusUnauthorized, "empty auth header")
		return
	}

	splitHeader := strings.Split(header, " ")
	if len(splitHeader) != 2 || splitHeader[0] != "Bearer" {
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
		return 0, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusUnauthorized,
			Msg:        "user id not found",
		}
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, &httpError.ErrorWithStatusCode{
			HTTPStatus: http.StatusUnauthorized,
			Msg:        "user id is of invalid type",
		}
	}

	return idInt, nil
}
