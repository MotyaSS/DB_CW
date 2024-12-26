package handler

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/MotyaSS/DB_CW/pkg/httpError"
	"github.com/gin-gonic/gin"
)

// Function aborts with status code and message
func abortWithStatusCode(ctx *gin.Context, statusCode int, msg string) {
	slog.Debug("abortWithStatusCode invoked",
		"statusCode", statusCode,
		"msg", msg)
	ctx.AbortWithStatusJSON(statusCode, map[string]string{"message": msg})
}

// Function tries to cast error to ErrorWithStatusCode and aborts with status code from it
// If error is not ErrorWithStatusCode, aborts with http.StatusInternalServerError
func abortWithError(ctx *gin.Context, err error) {
	slog.Debug("abortWithError invoked", "msg", err.Error())
	httpErr := &httpError.ErrorWithStatusCode{}
	ok := errors.As(err, &httpErr)
	if ok {
		abortWithStatusCode(ctx, httpErr.HTTPStatus, httpErr.Msg)
	} else {
		abortWithStatusCode(ctx, http.StatusInternalServerError, err.Error())
	}
	return
}
