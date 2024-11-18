package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (*Handler) getAllInstruments(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, "all instruments")
}

func (*Handler) getInstrument(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "instrument number "+ctx.Param("inst_id"))
}

func (*Handler) addInstrument(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, "instrument added")
}

func (*Handler) deleteInstrument(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "instrument deleted")
}
