package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (*Handler) getAllStores(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		"all stores",
	)
}

func (*Handler) getStore(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		[]string{
			"store ",
			ctx.Param("store_id"),
		},
	)
}

func (*Handler) createStore(ctx *gin.Context) {
	ctx.JSON(
		http.StatusCreated,
		"store created",
	)
}

func (*Handler) deleteStore(ctx *gin.Context) {
	ctx.JSON(
		http.StatusCreated,
		[]string{
			"store deleted",
			ctx.Param("store_id"),
		},
	)
}
