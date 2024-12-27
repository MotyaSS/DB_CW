package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllStores(ctx *gin.Context) {
	stores, err := h.service.Store.GetAllStores()

	if err != nil {
		abortWithError(ctx, err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		stores,
	)
}

func (h *Handler) getStore(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		[]string{
			"store ",
			ctx.Param("store_id"),
		},
	)
}

func (h *Handler) createStore(ctx *gin.Context) {
	ctx.JSON(
		http.StatusCreated,
		"store created",
	)
}

func (h *Handler) deleteStore(ctx *gin.Context) {
	ctx.JSON(
		http.StatusCreated,
		[]string{
			"store deleted",
			ctx.Param("store_id"),
		},
	)
}
