package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, "hi there")
	})
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", signUp)
		auth.POST("/sign-in", signIn)
	}
	items := router.Group("/instruments")
	{
		items.GET("/", getAllInstruments)
		items.POST("/new", addInstrument)
		item := items.Group("/:inst_id")
		{
			item.GET("/", getInstrument)
			item.DELETE("/", deleteInstrument)
			reviews := item.Group("/reviews")
			{
				reviews.GET("/", getReviews)
				reviews.POST("/", createReview)
				reviews.DELETE("/:review_id", deleteReview)
			}
		}
	}
	return router
}
