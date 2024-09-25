package services

import (
	"errors"

	"github.com/iki-rumondor/go-p3k/internal/app/layers/interfaces"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/models"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/request"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/response"
	"gorm.io/gorm"
)

type TransactionService struct {
	Repo interfaces.TransactionInterface
}

func NewTransactionService(repo interfaces.TransactionInterface) *TransactionService {
	return &TransactionService{
		Repo: repo,
	}
}

func (s *TransactionService) BuyProduct(userUuid string, req *request.BuyProduct) error {

	product, err := s.Repo.GetProductByUuid(req.ProductUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Produk tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}

	user, err := s.Repo.GetUserByUuid(userUuid)
	if err != nil {
		return response.SERVICE_INTERR
	}

	if user.Role.Name != "GUEST" && user.Role.Name != "CITIZEN" {
		return response.BADREQ_ERR("Silahkan gunakan akun pembeli atau akun masyarakat untuk transaksi produk")
	}

	if product.Stock < req.Quantity {
		return response.BADREQ_ERR("Jumlah pembelian melebihi stok produk")
	}

	model := models.ProductTransaction{
		ProductID: product.ID,
		Quantity:  req.Quantity,
	}

	if err := s.Repo.BuyProduct(userUuid, &model); err != nil {
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *TransactionService) DeleteProductTransaction(userUuid, transactionUuid string) error {

	if isHas := s.Repo.CheckProductTransactionIsAccept(transactionUuid); isHas {
		return response.BADREQ_ERR("Transaksi produk sudah disetujui")
	}

	if err := s.Repo.DeleteProductTransaction(userUuid, transactionUuid); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Transaksi produk tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *TransactionService) AcceptProductTransaction(userUuid, transactionUuid string) error {

	transaction, err := s.Repo.GetOwnerProductTransactionByUuid(userUuid, transactionUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Transaksi tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}

	if transaction.IsResponse {
		return response.NOTFOUND_ERR("Transaksi sudah direspon sebelumnya")
	}

	if transaction.Quantity > transaction.Product.Stock {
		return response.BADREQ_ERR("Stok tidak mencukupi")
	}

	model := models.ProductTransaction{
		ID:         transaction.ID,
		ProductID:  transaction.ProductID,
		Quantity:   transaction.Quantity,
		IsResponse: true,
		IsAccept:   true,
	}

	if err := s.Repo.AcceptProductTransaction(&model); err != nil {
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *TransactionService) UnacceptProductTransaction(userUuid, transactionUuid string) error {

	transaction, err := s.Repo.GetOwnerProductTransactionByUuid(userUuid, transactionUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Transaksi tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}

	if transaction.IsResponse {
		return response.NOTFOUND_ERR("Transaksi sudah direspon sebelumnya")
	}

	model := models.ProductTransaction{
		ID:         transaction.ID,
		IsResponse: true,
	}

	if err := s.Repo.UpdateModel(&model); err != nil {
		return response.SERVICE_INTERR
	}

	return nil
}
