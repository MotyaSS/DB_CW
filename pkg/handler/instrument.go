package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (*Handler) getAllInstruments(ctx *gin.Context) {
	categ, exists := ctx.GetQuery("category")
	response := "all instruments"
	if exists {
		response += " of category " + categ
	}
	ctx.JSON(
		http.StatusOK,
		response,
	)
}

func (*Handler) getInstrument(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		[]string{
			"instrument number ",
			ctx.Param("inst_id"),
		},
	)
}

func (*Handler) addInstrument(ctx *gin.Context) {
	ctx.JSON(
		http.StatusCreated,
		"instrument added",
	)
}

func (*Handler) deleteInstrument(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		[]string{
			"instrument deleted",
			ctx.Param("inst_id"),
		},
	)
}

func (*Handler) rentInstrument(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		"instrument rented",
	)
}
