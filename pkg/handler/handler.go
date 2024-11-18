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
		items.POST("/new", h.addInstrument)
		item := items.Group("/:inst_id")
		{
			item.GET("/", h.getInstrument)
			item.DELETE("/", h.deleteInstrument)
			reviews := item.Group("/reviews")
			{
				reviews.GET("/", h.getReviews)
				reviews.POST("/", h.createReview)
				reviews.DELETE("/:review_id", h.deleteReview)
			}
			rent := item.Group("/rent")
			{
				rent.POST("/")
			}
		}
	}
	stores := router.Group("/store")
	{
		stores.GET("/")
		stores.POST("/")
		stores.GET("/:store_id")
		stores.DELETE("/:store_id")
	}
	users := router.Group("/user")
	{
		users.GET("/")
		users.GET("/:user_id")
		users.DELETE("/:user_id")
	}
	return router
}
