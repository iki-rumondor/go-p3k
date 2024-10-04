package interfaces

import "github.com/iki-rumondor/go-p3k/internal/app/structs/models"

type TransactionInterface interface {
	CheckProductTransactionIsAccept(transactionUuid string) bool
	CheckProductTransactionIsResponse(transactionUuid string) bool
	GetOwnerProductTransactionByUuid(userUuid, transactionUuid string) (*models.ProductTransaction, error)
	GetProductByUuid(productUuid string) (*models.Product, error)
	GetUserByUuid(userUuid string) (*models.User, error)

	BuyProduct(userUuid string, model *models.ProductTransaction) error
	AcceptProductTransaction(model *models.ProductTransaction) error
	DeleteProductTransaction(userUuid, transactionUuid string) error

	UpdateModel(modelPointer interface{}) error
	UpdateTransaction(userUuid, uuid string, model *models.ProductTransaction) error
}
