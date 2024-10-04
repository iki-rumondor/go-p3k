package repositories

import (
	"github.com/iki-rumondor/go-p3k/internal/app/layers/interfaces"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/models"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/response"
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

func (r *TransactionRepo) CheckProductTransactionIsResponse(transactionUuid string) bool {
	rows := r.db.First(&models.ProductTransaction{}, "uuid = ? AND is_response = ?", transactionUuid, true).RowsAffected
	return rows == 1
}

func (r *TransactionRepo) GetProductByUuid(productUuid string) (*models.Product, error) {

	var product models.Product
	if err := r.db.First(&product, "uuid = ?", productUuid).Error; err != nil {
		return nil, err
	}

	return &product, nil

}

func (r *TransactionRepo) GetOwnerProductTransactionByUuid(userUuid, transactionUuid string) (*models.ProductTransaction, error) {
	var user models.User
	if err := r.db.Preload("Shop").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	var productIDs []uint
	if err := r.db.Model(&models.Product{}).Where("shop_id = ?", user.Shop.ID).Pluck("id", &productIDs).Error; err != nil {
		return nil, err
	}

	var transaction models.ProductTransaction
	if err := r.db.Preload("Product").First(&transaction, "uuid = ? AND product_id IN (?)", transactionUuid, productIDs).Error; err != nil {
		return nil, err
	}

	return &transaction, nil

}
func (r *TransactionRepo) GetUserByUuid(userUuid string) (*models.User, error) {

	var user models.User
	if err := r.db.Preload("Role").Preload("Member").First(&user, "uuid = ?", userUuid).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *TransactionRepo) GetActivityByUuid(uuid string) (*models.Activity, error) {
	var data models.Activity
	if err := r.db.First(&data, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &data, nil
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

func (r *TransactionRepo) AcceptProductTransaction(model *models.ProductTransaction) error {
	var product models.Product
	if err := r.db.First(&product, "id = ?", model.ProductID).Error; err != nil {
		return err
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		updatedProduct := models.Product{
			ID: product.ID,
		}

		if err := tx.Model(&updatedProduct).Update("stock", product.Stock-model.Quantity).Error; err != nil {
			return err
		}

		return tx.Updates(model).Error
	})

}

func (r *TransactionRepo) UpdateModel(modelPointer interface{}) error {
	return r.db.Updates(modelPointer).Error
}

func (r *TransactionRepo) UpdateTransaction(userUuid, uuid string, model *models.ProductTransaction) error {
	var user models.User
	if err := r.db.First(&user, "uuid = ?", userUuid).Error; err != nil {
		return err
	}

	var transaction models.ProductTransaction
	if err := r.db.First(&transaction, "uuid = ? AND user_id = ?", uuid, user.ID).Error; err != nil {
		return err
	}

	if transaction.ProofFile != "" {
		return response.BADREQ_ERR("Bukti transaksi sudah diupload")
	}

	model.ID = transaction.ID
	return r.db.Updates(model).Error
}

func (r *TransactionRepo) CreateModel(pointerModel interface{}) error {
	return r.db.Create(pointerModel).Error
}
