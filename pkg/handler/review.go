package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (*Handler) getAllReviews(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		[]string{
			"all reviews for instrument ",
			ctx.Param("inst_id"),
		},
	)
}

func (*Handler) getReview(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		[]string{
			"review",
			ctx.Param("review_id"),
			"for instrument",
			ctx.Param("inst_id"),
		},
	)
}

func (*Handler) createReview(ctx *gin.Context) {
	ctx.JSON(
		http.StatusCreated,
		"review created")
}

func (*Handler) deleteReview(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		[]string{
			"review deleted",
			ctx.Param(":review_id"),
		},
	)
}
