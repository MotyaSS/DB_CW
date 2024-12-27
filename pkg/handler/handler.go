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

func (h *Handler) InitRouter(middleware ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middleware...)
	apiRouter := router.Group("/api")
	apiRouter.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome page")
	})


	auth := apiRouter.Group("/")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/roles", h.getAllRoles)
		auth.POST("/sign-up-privileged", h.userIdentity, h.signUpPrivileged)
	}

	items := apiRouter.Group("/instruments")
	{
		items.GET("/", h.getAllInstruments)
		items.POST("/", h.userIdentity, h.createInstrument)

		item := items.Group("/:inst_id")
		{
			item.POST("/rent", h.userIdentity, h.rentInstrument)
			item.GET("/", h.getInstrument)
			item.DELETE("/", h.userIdentity, h.deleteInstrument)
			repairment := item.Group("/repairments")
			{
				repairment.POST("/", h.userIdentity, h.createRepair)
				repairment.GET("/", h.getAllRepairs)
				repairment.GET("/:repairment_id", h.getRepair)
			}
			reviews := item.Group("/reviews")
			{
				reviews.DELETE("/:review_id", h.userIdentity, h.deleteReview)
				reviews.GET("/:review_id", h.getReview)
				reviews.GET("/", h.getAllReviews)
				reviews.POST("/", h.userIdentity, h.createReview)
			}
		}
	}

	stores := apiRouter.Group("/stores")
	{
		stores.GET("/", h.getAllStores)
		stores.GET("/:store_id", h.getStore)
		stores.POST("/:store_id", h.userIdentity, h.createStore)
		stores.DELETE("/:store_id", h.userIdentity, h.deleteStore)
	}
	users := apiRouter.Group("/users", h.userIdentity)
	{
		users.GET("/", h.getAllUsers)
		users.GET("/:user_id", h.getUser)
		users.DELETE("/:user_id", h.deleteUser)
	}
	return router
}
