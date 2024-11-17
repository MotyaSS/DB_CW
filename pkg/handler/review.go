package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getReviews(ctx *gin.Context) {
	ctx.JSON(http.StatusOK,
		"all reviews for instrument "+
			ctx.Param("inst_id"))
}

func createReview(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated,
		"review created")
}

func deleteReview(ctx *gin.Context) {
	ctx.JSON(http.StatusOK,
		"review"+
			ctx.Param(":review_id")+
			"deleted")
}
