package repositories

import (
	"github.com/iki-rumondor/go-p3k/internal/app/layers/interfaces"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/models"
	"gorm.io/gorm"
)

type TransactionRepo struct {
	db *gorm.DB
}

func NewTransactionInterface(db *gorm.DB) interfaces.TransactionInterface {
	return &TransactionRepo{
		db: db,
	}
}

func (r *TransactionRepo) CheckProductTransactionIsAccept(transactionUuid string) bool {
	rows := r.db.First(&models.ProductTransaction{}, "uuid = ? AND is_accept = ?", transactionUuid, true).RowsAffected
	return rows == 1
}

func (r *TransactionRepo) GetProductByUuid(productUuid string) (*models.Product, error) {

	var product models.Product
	if err := r.db.First(&product, "uuid = ?", productUuid).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *TransactionRepo) BuyProduct(userUuid string, model *models.ProductTransaction) error {
	var user models.User
	if err := r.db.First(&user, "uuid = ?", userUuid).Error; err != nil {
		return err
	}

	model.UserID = user.ID

	return r.db.Create(model).Error
}

func (r *TransactionRepo) DeleteProductTransaction(userUuid, transactionUuid string) error {
	var user models.User
	if err := r.db.First(&user, "uuid = ?", userUuid).Error; err != nil {
		return err
	}

	var transaction models.ProductTransaction
	if err := r.db.First(&transaction, "uuid = ? AND user_id = ?", transactionUuid, user.ID).Error; err != nil {
		return err
	}

	return r.db.Delete(&transaction).Error
}
