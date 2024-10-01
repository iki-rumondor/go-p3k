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
		public.POST("/register/shop", handlers.AuthHandler.RegisterShop)

		public.GET("/public/products", handlers.FetchHandler.GetAllProducts)
		public.GET("/public/products/:uuid", handlers.FetchHandler.GetPublicProductByUuid)

		public.GET("/activities", handlers.FetchHandler.GetActivities)
		public.GET("/activities/:uuid", handlers.FetchHandler.GetActivityByUuid)

		public.GET("/files/products/:filename", handlers.FetchHandler.GetProductImage)
		public.GET("/files/activities/:filename", handlers.FetchHandler.GetActivityImage)
	}

	user := router.Group("api").Use(IsValidJWT()).Use(SetUserUuid())
	{
		user.GET("/users/detail", handlers.AuthHandler.GetUserByUuid)

		user.PATCH("/products/buy", handlers.TransactionHandler.BuyProduct)
		user.GET("/transactions", handlers.FetchHandler.GetProductTransactions)
		user.DELETE("/transactions/:uuid", handlers.TransactionHandler.DeleteProductTransaction)

		user.POST("/activities", handlers.ManagementHandler.CreateActivity)
		user.PUT("/activities/:uuid", handlers.ManagementHandler.UpdateActivity)
		user.DELETE("/activities/:uuid", handlers.ManagementHandler.DeleteActivity)

		user.GET("/members", handlers.FetchHandler.GetMembers)
		user.GET("/members/not/activities/:activityUuid", handlers.FetchHandler.GetMembersNotInActivity)
		user.POST("/members/activities", handlers.ManagementHandler.CreateMemberActivity)
		user.DELETE("/members/:memberUuid/activities/:activityUuid", handlers.ManagementHandler.DeleteMemberActivity)
	}

	admin := router.Group("api").Use(IsValidJWT()).Use(IsRole("ADMIN"))
	{
		admin.GET("/guests", handlers.FetchHandler.GetGuests)
		admin.GET("/guests/:uuid", handlers.FetchHandler.GetGuestByUuid)
		admin.PATCH("/users/activation/:uuid", handlers.AuthHandler.ActivationUser)

		admin.GET("/citizens", handlers.FetchHandler.GetCitizens)
		admin.GET("/citizens/:uuid", handlers.FetchHandler.GetCitizenByUuid)
		admin.POST("/citizens", handlers.ManagementHandler.CreateCitizen)
		admin.PUT("/citizens/:uuid", handlers.ManagementHandler.UpdateCitizen)

		admin.GET("/categories", handlers.FetchHandler.GetCategories)
		admin.GET("/categories/:uuid", handlers.FetchHandler.GetCategoryByUuid)
		admin.POST("/categories", handlers.ManagementHandler.CreateCategory)
		admin.PUT("/categories/:uuid", handlers.ManagementHandler.UpdateCategory)

		admin.GET("/shops", handlers.FetchHandler.GetShops)
		admin.GET("/shops/:uuid", handlers.FetchHandler.GetShopByUuid)
		// admin.POST("/shops", handlers.ManagementHandler.CreateShop)
		// admin.PUT("/shops/:uuid", handlers.ManagementHandler.UpdateShop)

		admin.GET("/members/:uuid", handlers.FetchHandler.GetMemberByUuid)
		admin.POST("/members", handlers.ManagementHandler.CreateMember)
		admin.PUT("/members/:uuid", handlers.ManagementHandler.UpdateMember)

	}

	umkm := router.Group("api").Use(IsValidJWT()).Use(IsRole("UMKM")).Use(SetUserUuid())
	{
		umkm.GET("/products", handlers.FetchHandler.GetProducts)
		umkm.GET("/products/:uuid", handlers.FetchHandler.GetProductByUuid)
		umkm.POST("/products", handlers.ManagementHandler.CreateProduct)
		umkm.PUT("/products/:uuid", handlers.ManagementHandler.UpdateProduct)

		umkm.GET("/shops/transactions", handlers.FetchHandler.GetProductTransactionsByShop)
		umkm.GET("/transactions/:uuid", handlers.FetchHandler.GetProductTransactionByUuid)
		umkm.PATCH("/transactions/:transactionUuid/accept", handlers.TransactionHandler.AcceptProductTransaction)
		umkm.PATCH("/transactions/:transactionUuid/unaccept", handlers.TransactionHandler.UnacceptProductTransaction)
	}

	return router
}
