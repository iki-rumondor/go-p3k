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
	// 1. Cancel transactions where seller has not responded within 10 minutes from created_at
	tenMinutesAgoCreated := time.Now().Add(-10 * time.Minute).UnixMilli()
	var unconfirmed []models.ProductTransaction
	err := db.Preload("Product").Where("is_response = ? AND created_at < ?", false, tenMinutesAgoCreated).Find(&unconfirmed).Error
	if err != nil {
		log.Println("Scheduler error autoCancelUnpaidTransactions (unconfirmed):", err.Error())
	} else {
		for _, tx := range unconfirmed {
			cancelTransaction(db, tx, "otomatis dibatalkan karena penjual tidak merespon dalam 10 menit")
		}
	}

	// 2. Cancel transactions where seller accepted, but buyer hasn't uploaded payment proof within 10 minutes from updated_at
	tenMinutesAgoUpdated := time.Now().Add(-10 * time.Minute).UnixMilli()
	var unpaid []models.ProductTransaction
	err = db.Preload("Product").Where("is_response = ? AND is_accept = ? AND proof_file = ? AND payment_verified = ? AND updated_at < ?", true, true, "", false, tenMinutesAgoUpdated).Find(&unpaid).Error
	if err != nil {
		log.Println("Scheduler error autoCancelUnpaidTransactions (unpaid):", err.Error())
	} else {
		for _, tx := range unpaid {
			cancelTransaction(db, tx, "otomatis dibatalkan karena pembeli tidak membayar dalam 10 menit setelah disetujui penjual")
		}
	}
}

func cancelTransaction(db *gorm.DB, tx models.ProductTransaction, reason string) {
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

		log.Printf("Scheduler: Transaksi %s %s. Stok produk kembali ditambahkan.\n", tx.Uuid, reason)
		return nil
	})
	if err != nil {
		log.Println("Scheduler error executing tx cancel:", err.Error())
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
