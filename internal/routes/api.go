package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-p3k/internal/config"
)

func StartServer(handlers *config.Handlers) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"http://localhost:5173", "http://103.26.13.166:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12,
	}))

	public := router.Group("api")
	{
		public.POST("/verify-user", handlers.AuthHandler.VerifyUser)
	}

	return router
}
