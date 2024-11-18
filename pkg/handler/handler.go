package handler

import (
	"net/http"

	"github.com/MotyaSS/DB_CW/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome page")
	})
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	items := router.Group("/instruments")
	{
		items.GET("/", h.getAllInstruments)
		items.POST("/", h.addInstrument)
		item := items.Group("/:inst_id")
		{
			item.GET("/", h.getInstrument)
			item.DELETE("/", h.deleteInstrument)
			item.POST("/rent", h.rentInstrument)
			reviews := item.Group("/reviews")
			{
				reviews.DELETE("/:review_id", h.deleteReview)
				reviews.GET("/:review_id", h.getReview)
				reviews.GET("/", h.getReviews)
				reviews.POST("/", h.createReview)
			}
		}
	}
	stores := router.Group("/store")
	{
		stores.GET("/", h.getAllStores)
		stores.GET("/:store_id", h.getStore)
		stores.POST("/:store_id", h.createStore)
		stores.DELETE("/:store_id", h.deleteStore)
	}
	users := router.Group("/user")
	{
		users.GET("/", h.getUsers)
		users.GET("/:user_id", h.getUser)
		users.DELETE("/:user_id", h.deleteUser)
	}
	return router
}
