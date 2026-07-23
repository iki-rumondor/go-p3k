package services

import (
	"errors"
	"log"
	"time"

	"github.com/iki-rumondor/go-p3k/internal/app/layers/interfaces"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/models"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/request"
	"github.com/iki-rumondor/go-p3k/internal/app/structs/response"
	"github.com/iki-rumondor/go-p3k/internal/utils"
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

	model := models.ProductTransaction{
		ID:         transaction.ID,
		ProductID:  transaction.ProductID,
		Quantity:   transaction.Quantity,
		IsResponse: true,
		IsAccept:   true,
		Revenue:    transaction.Quantity * transaction.Product.Price,
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
		return response.BADREQ_ERR("Transaksi sudah direspon sebelumnya")
	}

	if transaction.ProofFile != "" {
		return response.BADREQ_ERR("Tindakan tidak dapat dilanjutkan, bukti transaksi sudah diupload")
	}

	model := models.ProductTransaction{
		ID:         transaction.ID,
		ProductID:  transaction.ProductID,
		Quantity:   transaction.Quantity,
		IsResponse: true,
		IsAccept:   false,
	}

	if err := s.Repo.UnacceptProductTransaction(&model); err != nil {
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *TransactionService) ConfirmProductTransaction(userUuid, transactionUuid string) error {
	transaction, err := s.Repo.GetOwnerProductTransactionByUuid(userUuid, transactionUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Transaksi tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}

	if transaction.IsResponse {
		return response.BADREQ_ERR("Transaksi sudah direspon sebelumnya")
	}

	if transaction.IsConfirm {
		return response.BADREQ_ERR("Transaksi sudah dikonfirmasi sebelumnya")
	}

	model := models.ProductTransaction{
		ID:        transaction.ID,
		IsConfirm: true,
	}

	if err := s.Repo.UpdateModel(&model); err != nil {
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *TransactionService) SetTransactionProof(userUuid, transactionUuid, filename string) error {
	transaction, err := s.Repo.GetTransactionByUuid(transactionUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Transaksi tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}

	if !transaction.IsResponse {
		return response.BADREQ_ERR("Transaksi belum disetujui oleh penjual")
	}

	if !transaction.IsAccept {
		return response.BADREQ_ERR("Transaksi telah ditolak oleh penjual")
	}

	if transaction.PaymentVerified {
		return response.BADREQ_ERR("Pembayaran transaksi sudah diverifikasi sebelumnya")
	}

	model := models.ProductTransaction{
		ProofFile: filename,
	}

	if err := s.Repo.UpdateTransaction(userUuid, transactionUuid, &model); err != nil {
		log.Println(err.Error())
		if utils.IsErrorType(err) {
			return err
		}
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *TransactionService) CreateMemberActivity(userUuid, activityUuid, filename string) error {
	user, err := s.Repo.GetUserByUuid(userUuid)
	if err != nil {
		return response.SERVICE_INTERR
	}

	activity, err := s.Repo.GetActivityByUuid(activityUuid)
	if err != nil {
		return response.SERVICE_INTERR
	}

	now := time.Now().UnixMilli()

	if now < activity.StartTime {
		return response.BADREQ_ERR("Kegiatan belum dimulai")
	}

	if now > activity.EndTime {
		return response.BADREQ_ERR("Kegiatan sudah selesai")
	}

	model := models.MemberActivity{
		MemberID:        user.Member.ID,
		ActivityID:      activity.ID,
		AttendenceImage: filename,
	}

	if err := s.Repo.CreateModel(&model); err != nil {
		log.Println(err.Error())
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *TransactionService) VerifyPayment(userUuid string, transactionUuid string) error {
	transaction, err := s.Repo.GetTransactionByUuid(transactionUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Transaksi tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}

	if transaction.PaymentVerified {
		return response.BADREQ_ERR("Pembayaran transaksi sudah diverifikasi sebelumnya")
	}

	if transaction.ProofFile == "" {
		return response.BADREQ_ERR("Pembeli belum mengunggah bukti pembayaran")
	}

	model := models.ProductTransaction{
		ID:              transaction.ID,
		PaymentVerified: true,
		IsConfirm:       true,
	}

	if err := s.Repo.UpdateModel(&model); err != nil {
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *TransactionService) RejectPayment(userUuid string, transactionUuid string) error {
	transaction, err := s.Repo.GetTransactionByUuid(transactionUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Transaksi tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}

	if transaction.PaymentVerified {
		return response.BADREQ_ERR("Pembayaran sudah diverifikasi, tidak dapat menolak")
	}

	model := models.ProductTransaction{
		ID:         transaction.ID,
		ProductID:  transaction.ProductID,
		Quantity:   transaction.Quantity,
		IsResponse: true,
		IsAccept:   false,
	}

	if err := s.Repo.UnacceptProductTransaction(&model); err != nil {
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *TransactionService) ConfirmDelivery(userUuid string, transactionUuid string, filename string) error {
	transaction, err := s.Repo.GetOwnerProductTransactionByUuid(userUuid, transactionUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Transaksi tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}

	if !transaction.PaymentVerified {
		return response.BADREQ_ERR("Pembayaran transaksi belum diverifikasi oleh admin")
	}

	if transaction.IsDelivered {
		return response.BADREQ_ERR("Barang sudah dikonfirmasi terkirim sebelumnya")
	}

	model := models.ProductTransaction{
		ID:            transaction.ID,
		IsDelivered:   true,
		DeliveredAt:   time.Now().UnixMilli(),
		DeliveryProof: filename,
	}

	if err := s.Repo.UpdateModel(&model); err != nil {
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *TransactionService) ConfirmReceipt(userUuid string, transactionUuid string) error {
	transaction, err := s.Repo.GetTransactionByUuid(transactionUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Transaksi tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}

	user, err := s.Repo.GetUserByUuid(userUuid)
	if err != nil {
		return response.SERVICE_INTERR
	}

	if transaction.UserID != user.ID {
		return response.BADREQ_ERR("Anda tidak memiliki akses ke transaksi ini")
	}

	if !transaction.PaymentVerified {
		return response.BADREQ_ERR("Pembayaran transaksi belum diverifikasi oleh admin")
	}

	if transaction.IsDelivered {
		return response.BADREQ_ERR("Barang sudah diterima/dikonfirmasi sebelumnya")
	}

	model := models.ProductTransaction{
		ID:          transaction.ID,
		IsDelivered: true,
		DeliveredAt: time.Now().UnixMilli(),
	}

	if err := s.Repo.UpdateModel(&model); err != nil {
		return response.SERVICE_INTERR
	}

	return nil
}

func (s *TransactionService) Disburse(userUuid string, transactionUuid string) error {
	transaction, err := s.Repo.GetTransactionByUuid(transactionUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NOTFOUND_ERR("Transaksi tidak ditemukan")
		}
		return response.SERVICE_INTERR
	}

	if !transaction.PaymentVerified {
		return response.BADREQ_ERR("Pembayaran transaksi belum diverifikasi")
	}

	if !transaction.IsDelivered {
		return response.BADREQ_ERR("Barang belum dikirim/diterima oleh pembeli")
	}

	if transaction.IsDisbursed {
		return response.BADREQ_ERR("Dana transaksi sudah disalurkan sebelumnya")
	}

	model := models.ProductTransaction{
		ID:          transaction.ID,
		IsDisbursed: true,
		IsResponse:  true,
		IsAccept:    true,
		Revenue:     transaction.Quantity * transaction.Product.Price,
	}

	if err := s.Repo.UpdateModel(&model); err != nil {
		return response.SERVICE_INTERR
	}

	return nil
}

