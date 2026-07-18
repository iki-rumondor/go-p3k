package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-p3k/internal/app/layers/newest"
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
		public.GET("/files/attendances/:filename", handlers.FetchHandler.GetAttendanceImage)
		public.GET("/shops", handlers.FetchHandler.GetShops)
		public.GET("/files/shops/:filename", handlers.FetchHandler.GetShopImage)
		public.GET("/files/qris/:filename", handlers.FetchHandler.GetQrisImage)
		public.GET("/public/admin/qris", handlers.FetchHandler.GetAdminQris)
	}

	user := router.Group("api").Use(IsValidJWT()).Use(SetUserUuid())
	{
		user.GET("/users/detail", handlers.AuthHandler.GetUserByUuid)
		user.GET("/categories", handlers.FetchHandler.GetCategories)

		user.PATCH("/products/buy", handlers.TransactionHandler.BuyProduct)
		user.GET("/transactions", handlers.FetchHandler.GetProductTransactions)
		user.DELETE("/transactions/:transactionUuid", handlers.TransactionHandler.DeleteProductTransaction)

		user.POST("/activities", handlers.ManagementHandler.CreateActivity)
		user.PUT("/activities/:uuid", handlers.ManagementHandler.UpdateActivity)
		user.DELETE("/activities/:uuid", handlers.ManagementHandler.DeleteActivity)

		user.GET("/members", handlers.FetchHandler.GetMembers)
		user.GET("/guests", handlers.FetchHandler.GetGuests)
		user.GET("/citizens", handlers.FetchHandler.GetCitizens)

		user.GET("/members/not/activities/:activityUuid", handlers.FetchHandler.GetMembersNotInActivity)
		user.POST("/members/activities", handlers.ManagementHandler.CreateMemberActivity)
		user.DELETE("/members/:memberUuid/activities/:activityUuid", handlers.ManagementHandler.DeleteMemberActivity)

		user.PATCH("/transactions/:transactionUuid/proof", handlers.TransactionHandler.SetTransactionProof)
		user.GET("/files/transaction_proofs/:filename", handlers.FetchHandler.GetTransactionProofImage)
		user.GET("/files/delivery_proofs/:filename", handlers.FetchHandler.GetDeliveryProofImage)

		user.PATCH("/users/password", handlers.AuthHandler.UpdatePassword)

		user.GET("/dashboard/guest", handlers.FetchHandler.GetGuestDashboard)
		user.PUT("/guests/:uuid", handlers.ManagementHandler.UpdateGuest)

		user.GET("/tutorials", handlers.FetchHandler.GetTutorials)
		user.GET("/tutorials/:uuid", handlers.FetchHandler.GetTutorialByUuid)
		user.PATCH("/transactions/:transactionUuid/confirm-receipt", handlers.TransactionHandler.ConfirmReceipt)
	}

	admin := router.Group("api").Use(IsValidJWT()).Use(IsRole("ADMIN"))
	{
		admin.GET("/guests/:uuid", handlers.FetchHandler.GetGuestByUuid)
		admin.PATCH("/users/activation/:uuid", handlers.AuthHandler.ActivationUser)

		admin.GET("/citizens/:uuid", handlers.FetchHandler.GetCitizenByUuid)
		admin.POST("/citizens", handlers.ManagementHandler.CreateCitizen)
		admin.PUT("/citizens/:uuid", handlers.ManagementHandler.UpdateCitizen)
		admin.DELETE("/citizens/:uuid", handlers.ManagementHandler.DeleteCitizen)

		public.GET("/regions", newest.GetAllRegions)
		public.GET("/regions/:id", newest.GetRegionByID)
		admin.POST("/regions", newest.CreateRegion)
		admin.PUT("/regions/:id", newest.UpdateRegion)
		admin.DELETE("/regions/:id", newest.DeleteRegion)

		admin.GET("/categories/:uuid", handlers.FetchHandler.GetCategoryByUuid)
		admin.POST("/categories", handlers.ManagementHandler.CreateCategory)
		admin.PUT("/categories/:uuid", handlers.ManagementHandler.UpdateCategory)
		admin.DELETE("/categories/:uuid", handlers.ManagementHandler.DeleteCategory)

		admin.GET("/shops/:uuid", handlers.FetchHandler.GetShopByUuid)

		admin.GET("/members/:uuid", handlers.FetchHandler.GetMemberByUuid)
		admin.POST("/members", handlers.ManagementHandler.CreateMember)
		admin.PUT("/members/:uuid", handlers.ManagementHandler.UpdateMember)
		admin.DELETE("/master/members/:uuid", handlers.ManagementHandler.DeleteMember)

		admin.GET("/files/identities/:filename", handlers.FetchHandler.GetIdentityImage)

		admin.GET("/dashboard/admin", handlers.FetchHandler.GetAdminDashboard)
		admin.PATCH("/users/reset-password/:uuid", handlers.AuthHandler.ResetPassword)
		admin.GET("/member-activities", handlers.FetchHandler.GetAllMemberActivities)
		admin.PATCH("/member-activities/:uuid/accept-attendance", handlers.ManagementHandler.UpdatePresence)

		admin.GET("/admin/tutorials", handlers.FetchHandler.GetTutorials)
		admin.GET("/admin/tutorials/:uuid", handlers.FetchHandler.GetTutorialByUuid)
		admin.POST("/admin/tutorials", handlers.ManagementHandler.CreateTutorial)
		admin.PUT("/admin/tutorials/:uuid", handlers.ManagementHandler.UpdateTutorial)
		admin.DELETE("/admin/tutorials/:uuid", handlers.ManagementHandler.DeleteTutorial)

		admin.POST("/tutorials", handlers.ManagementHandler.CreateTutorial)
		admin.PUT("/tutorials/:uuid", handlers.ManagementHandler.UpdateTutorial)
		admin.DELETE("/tutorials/:uuid", handlers.ManagementHandler.DeleteTutorial)
		admin.PATCH("/settings/qris", handlers.ManagementHandler.UploadAdminQris)

		admin.GET("/admin/transactions", handlers.FetchHandler.GetAdminTransactions)
		admin.GET("/admin/transactions/:transactionUuid", handlers.FetchHandler.GetAdminTransactionByUuid)
		admin.PATCH("/admin/transactions/:transactionUuid/verify-payment", handlers.TransactionHandler.VerifyPayment)
		admin.PATCH("/admin/transactions/:transactionUuid/reject-payment", handlers.TransactionHandler.RejectPayment)
		admin.PATCH("/admin/transactions/:transactionUuid/disburse", handlers.TransactionHandler.Disburse)
	}

	umkm := router.Group("api").Use(IsValidJWT()).Use(IsRole("UMKM")).Use(SetUserUuid())
	{
		umkm.PUT("/shops/:uuid", handlers.ManagementHandler.UpdateShop)
		umkm.PATCH("/shops/qris", handlers.ManagementHandler.UploadQris)
		umkm.GET("/shops/user", handlers.FetchHandler.GetShopByUser)
		umkm.GET("/dashboard/shop", handlers.FetchHandler.GetShopDashboard)

		umkm.GET("/products", handlers.FetchHandler.GetProducts)
		umkm.GET("/products/:uuid", handlers.FetchHandler.GetProductByUuid)
		umkm.POST("/products", handlers.ManagementHandler.CreateProduct)
		umkm.PUT("/products/:uuid", handlers.ManagementHandler.UpdateProduct)
		umkm.DELETE("/products/:uuid", handlers.ManagementHandler.DeleteProduct)

		umkm.GET("/shops/transactions", handlers.FetchHandler.GetProductTransactionsByShop)
		umkm.GET("/transactions/:transactionUuid", handlers.FetchHandler.GetProductTransactionByUuid)
		umkm.PATCH("/transactions/:transactionUuid/accept", handlers.TransactionHandler.AcceptProductTransaction)
		umkm.PATCH("/transactions/:transactionUuid/unaccept", handlers.TransactionHandler.UnacceptProductTransaction)
		umkm.PATCH("/transactions/:transactionUuid/confirm", handlers.TransactionHandler.ConfirmProductTransaction)
		umkm.PATCH("/transactions/:transactionUuid/confirm-delivery", handlers.TransactionHandler.ConfirmDelivery)
	}

	member := router.Group("api").Use(IsValidJWT()).Use(IsRole("MEMBER")).Use(SetUserUuid())
	{
		member.GET("/members/activities/:activityUuid", handlers.FetchHandler.GetMemberActivity)
		member.POST("/activities/:activityUuid/attendance", handlers.TransactionHandler.CreateMemberActivity)
		member.GET("/dashboard/member", handlers.FetchHandler.GetMemberDashboard)
		member.GET("/member/activities", handlers.FetchHandler.GetMemberActivities)
		member.GET("/member/user", handlers.FetchHandler.GetMemberByUser)
	}

	return router
}
