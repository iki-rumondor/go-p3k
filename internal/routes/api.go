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
		public.POST("/register/guest", handlers.AuthHandler.RegisterGuest)
		public.GET("/public/products", handlers.FetchHandler.GetAllProducts)
		public.GET("/public/products/:uuid", handlers.FetchHandler.GetPublicProductByUuid)
	}

	user := router.Group("api").Use(IsValidJWT()).Use(SetUserUuid())
	{
		user.GET("/users/detail", handlers.AuthHandler.GetUserByUuid)

		user.PATCH("/products/buy", handlers.TransactionHandler.BuyProduct)
		user.GET("/transactions", handlers.FetchHandler.GetProductTransactions)
		user.DELETE("/transactions/:uuid", handlers.TransactionHandler.DeleteProductTransaction)
	}

	admin := router.Group("api").Use(IsValidJWT()).Use(IsRole("ADMIN"))
	{
		admin.GET("/guests", handlers.FetchHandler.GetGuests)
		admin.GET("/guests/:uuid", handlers.FetchHandler.GetGuestByUuid)
		admin.PATCH("/users/activation/:uuid", handlers.AuthHandler.ActivationUser)

		admin.GET("/categories", handlers.FetchHandler.GetCategories)
		admin.GET("/categories/:uuid", handlers.FetchHandler.GetCategoryByUuid)
		admin.POST("/categories", handlers.ManagementHandler.CreateCategory)
		admin.PUT("/categories/:uuid", handlers.ManagementHandler.UpdateCategory)

		admin.GET("/shops", handlers.FetchHandler.GetShops)
		admin.GET("/shops/:uuid", handlers.FetchHandler.GetShopByUuid)
		admin.POST("/shops", handlers.ManagementHandler.CreateShop)
		admin.PUT("/shops/:uuid", handlers.ManagementHandler.UpdateShop)

	}

	umkm := router.Group("api").Use(IsValidJWT()).Use(IsRole("UMKM")).Use(SetUserUuid())
	{
		umkm.GET("/products", handlers.FetchHandler.GetProducts)
		umkm.GET("/products/:uuid", handlers.FetchHandler.GetProductByUuid)
		umkm.POST("/products", handlers.ManagementHandler.CreateProduct)
		umkm.PUT("/products/:uuid", handlers.ManagementHandler.UpdateProduct)
	}

	return router
}
