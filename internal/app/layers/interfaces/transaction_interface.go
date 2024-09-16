package interfaces

import "github.com/iki-rumondor/go-p3k/internal/app/structs/models"

type TransactionInterface interface {
	CheckProductTransactionIsAccept(transactionUuid string) bool
	GetProductByUuid(productUuid string) (*models.Product, error)
	BuyProduct(userUuid string, model *models.ProductTransaction) error
	DeleteProductTransaction(userUuid, transactionUuid string) error
}
