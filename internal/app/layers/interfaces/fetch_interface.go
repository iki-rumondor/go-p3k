package interfaces

import "github.com/iki-rumondor/go-p3k/internal/app/structs/models"

type FetchInterface interface {
	GetGuests() (*[]models.Guest, error)
	GetGuestByUuid(uuid string) (*models.Guest, error)
	GetGuestByUser(userUuid string) (*models.Guest, error)

	GetCitizens() (*[]models.Citizen, error)
	GetCitizenByUuid(uuid string) (*models.Citizen, error)

	GetCategories() (*[]models.Category, error)
	GetCategoryByUuid(uuid string) (*models.Category, error)

	GetShops(limit int) (*[]models.Shop, error)
	GetShopByUuid(uuid string) (*models.Shop, error)
	GetShopByUser(userUuid string) (*models.Shop, error)

	GetAllProducts(limit int, shopUuid string) (*[]models.Product, error)
	GetPublicProductByUuid(uuid string) (*models.Product, error)
	GetProducts(userUuid string) (*[]models.Product, error)
	GetProductByUuid(userUuid, uuid string) (*models.Product, error)

	GetProductTransactions(userUuid string) (*[]models.ProductTransaction, error)
	GetProductTransactionByUuid(userUuid, uuid string) (*models.ProductTransaction, error)
	GetProductTransactionsByShop(userUuid string, isAccept bool) (*[]models.ProductTransaction, error)

	GetMembers(group string) (*[]models.Member, error)
	GetMemberByUuid(uuid string) (*models.Member, error)
	GetMemberByUserUuid(uuid string) (*models.Member, error)

	GetActivitiesWithDate(limit int, group, startDate, endDate string) (*[]models.Activity, error)
	GetActivities(limit int, group string) (*[]models.Activity, error)
	GetActivityByUuid(uuid string) (*models.Activity, error)

	GetMembersNotInActivity(activityUuid string) (*[]models.Member, error)
	GetMemberActivity(userUuid, activityUuid string) (*models.MemberActivity, error)
	GetMemberActivities() (*[]models.MemberActivity, error)

	CountGuestsInactive() (int64, error)
	CountShopsInactive() (int64, error)

	CountShopProducts(userUuid string) (int64, error)
	CountShopUnprocessTransactions(userUuid string) (int64, error)

	CountUserSuccessTransactions(userUuid string) (int64, error)
	CountUserUnprocessTransactions(userUuid string) (int64, error)

	CountActivities() (int64, error)
}
