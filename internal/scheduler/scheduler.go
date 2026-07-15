package scheduler

import (
	"log"
	"time"

	"github.com/iki-rumondor/go-p3k/internal/app/structs/models"
	"gorm.io/gorm"
)

func StartScheduler(db *gorm.DB) {
	log.Println("Background scheduler started successfully.")
	// Goroutine for 10 minutes auto-cancel unpaid transactions
	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		for range ticker.C {
			autoCancelUnpaidTransactions(db)
		}
	}()

	// Goroutine for 3 days auto-completion of delivered transactions
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		for range ticker.C {
			autoCompleteDeliveredTransactions(db)
		}
	}()
}

func autoCancelUnpaidTransactions(db *gorm.DB) {
	var transactions []models.ProductTransaction
	tenMinutesAgo := time.Now().Add(-10 * time.Minute).UnixMilli()

	// Query transactions where payment is not verified, not responded yet, no proof uploaded, and older than 10 mins
	err := db.Preload("Product").Where("payment_verified = ? AND is_response = ? AND proof_file = ? AND created_at < ?", false, false, "", tenMinutesAgo).Find(&transactions).Error
	if err != nil {
		log.Println("Scheduler error autoCancelUnpaidTransactions:", err.Error())
		return
	}

	for _, tx := range transactions {
		err := db.Transaction(func(dbTx *gorm.DB) error {
			// Mark as rejected/cancelled
			if err := dbTx.Model(&tx).Updates(map[string]interface{}{
				"is_response": true,
				"is_accept":   false,
			}).Error; err != nil {
				return err
			}

			// Restore stock
			if tx.Product != nil {
				if err := dbTx.Model(&models.Product{ID: tx.ProductID}).Update("stock", tx.Product.Stock+tx.Quantity).Error; err != nil {
					return err
				}
			}

			log.Printf("Scheduler: Transaksi %s otomatis dibatalkan karena tidak dibayar dalam 10 menit. Stok produk kembali ditambahkan.\n", tx.Uuid)
			return nil
		})
		if err != nil {
			log.Println("Scheduler error executing tx cancel:", err.Error())
		}
	}
}

func autoCompleteDeliveredTransactions(db *gorm.DB) {
	var transactions []models.ProductTransaction
	threeDaysAgo := time.Now().Add(-3 * 24 * time.Hour).UnixMilli()

	// Query transactions where payment is verified, delivered by seller, but not disbursed (fully completed) by admin yet, and DeliveredAt is older than 3 days
	err := db.Preload("Product").Where("payment_verified = ? AND is_delivered = ? AND is_disbursed = ? AND delivered_at > 0 AND delivered_at < ?", true, true, false, threeDaysAgo).Find(&transactions).Error
	if err != nil {
		log.Println("Scheduler error autoCompleteDeliveredTransactions:", err.Error())
		return
	}

	for _, tx := range transactions {
		err := db.Transaction(func(dbTx *gorm.DB) error {
			// Mark as disbursed (completed)
			if err := dbTx.Model(&tx).Updates(map[string]interface{}{
				"is_disbursed": true,
				"is_response":  true,
				"is_accept":    true,
				"revenue":      tx.Quantity * tx.Product.Price,
			}).Error; err != nil {
				return err
			}

			log.Printf("Scheduler: Transaksi %s otomatis diselesaikan oleh sistem setelah 3 hari dikirim oleh penjual.\n", tx.Uuid)
			return nil
		})
		if err != nil {
			log.Println("Scheduler error executing tx disburse:", err.Error())
		}
	}
}
