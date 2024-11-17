package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllInstruments(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "all instruments")
}

func getInstrument(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "instrument number "+ctx.Param("inst_id"))
}

func addInstrument(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, "instrument added")
}

func deleteInstrument(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "instrument deleted")
}
