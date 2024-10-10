package interfaces

import "github.com/iki-rumondor/go-p3k/internal/app/structs/models"

type FetchInterface interface {
	GetGuests() (*[]models.Guest, error)
	GetGuestByUuid(uuid string) (*models.Guest, error)

	GetCitizens() (*[]models.Citizen, error)
	GetCitizenByUuid(uuid string) (*models.Citizen, error)

	GetCategories() (*[]models.Category, error)
	GetCategoryByUuid(uuid string) (*models.Category, error)

	GetShops() (*[]models.Shop, error)
	GetShopByUuid(uuid string) (*models.Shop, error)

	GetAllProducts(limit int) (*[]models.Product, error)
	GetPublicProductByUuid(uuid string) (*models.Product, error)
	GetProducts(userUuid string) (*[]models.Product, error)
	GetProductByUuid(userUuid, uuid string) (*models.Product, error)

	GetProductTransactions(userUuid string) (*[]models.ProductTransaction, error)
	GetProductTransactionByUuid(userUuid, uuid string) (*models.ProductTransaction, error)
	GetProductTransactionsByShop(userUuid string) (*[]models.ProductTransaction, error)

	GetMembers() (*[]models.Member, error)
	GetMemberByUuid(uuid string) (*models.Member, error)

	GetActivities(limit int) (*[]models.Activity, error)
	GetActivityByUuid(uuid string) (*models.Activity, error)

	GetMembersNotInActivity(activityUuid string) (*[]models.Member, error)
	GetMemberActivity(userUuid, activityUuid string) (*models.MemberActivity, error)

	CountGuestsInactive() (int64, error)
	CountShopsInactive() (int64, error)
}
