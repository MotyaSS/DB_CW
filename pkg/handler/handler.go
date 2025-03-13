package handler

import (
	"net/http"

	"github.com/MotyaSS/DB_CW/pkg/service"
	"github.com/gin-contrib/cors"
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

	// CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

	router.Use(cors.New(config))
	router.Use(middleware...)

	router.RedirectTrailingSlash = false
	router.RedirectFixedPath = false

	apiRouter := router.Group("/api")
	apiRouter.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome page")
	})

	auth := apiRouter.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/me", h.userIdentity, h.getCurrentUser)
		auth.GET("/roles", h.getAllRoles)
		auth.POST("/sign-up-privileged", h.userIdentity, h.signUpPrivileged)
	}

	items := apiRouter.Group("/instruments")
	{
		items.GET("", h.getAllInstruments)
		items.POST("", h.userIdentity, h.createInstrument)
		items.GET("/categories", h.getCategories)
		items.POST("/categories", h.userIdentity, h.createCategory)
		items.GET("/manufacturers", h.getManufacturers)
		items.POST("/manufacturers", h.userIdentity, h.createManufacturer)

		item := items.Group("/:inst_id")
		{
			item.GET("", h.getInstrument)
			item.DELETE("", h.userIdentity, h.deleteInstrument)
			repairment := item.Group("/repairments", h.userIdentity)
			{
				repairment.POST("", h.createRepair)
				repairment.GET("", h.getAllRepairs)
				repairment.GET("/:repairment_id", h.getRepair)
			}
			reviews := item.Group("/reviews")
			{
				reviews.DELETE("/:review_id", h.userIdentity, h.deleteReview)
				reviews.GET("/:review_id", h.getReview)
				reviews.GET("", h.getAllReviews)
				reviews.POST("", h.userIdentity, h.createReview)
			}
			rent := item.Group("/rent", h.userIdentity)
			{
				rent.POST("", h.createRental)
				rent.GET("", h.getInstrumentRentals)
			}
		}
	}

	stores := apiRouter.Group("/stores")
	{
		stores.GET("", h.getAllStores)
		stores.GET("/", h.getStore)
		stores.POST("/", h.userIdentity, h.createStore)
		stores.DELETE("/:store_id", h.userIdentity, h.deleteStore)
	}
	users := apiRouter.Group("/users", h.userIdentity)
	{
		users.GET("", h.getAllUsers)
		users.GET("/:user_id", h.getUser)
		users.DELETE("/:user_id", h.deleteUser)
	}

	rentals := apiRouter.Group("/rentals", h.userIdentity)
	{
		rentals.GET("", h.getUserRentals)
		rentals.GET("/:rental_id", h.getRental)
		rentals.POST("/:rental_id/return", h.returnInstrument)
	}

	// Добавляем новую группу маршрутов для бэкапов
	backups := apiRouter.Group("/backups", h.userIdentity)
	{
		backups.POST("", h.createBackup)
		backups.GET("", h.listBackups)
		backups.POST("/restore", h.restoreBackup)
	}

	return router
}
