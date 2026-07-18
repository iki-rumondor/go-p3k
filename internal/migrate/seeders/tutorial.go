package seeders

import (
	"github.com/iki-rumondor/go-p3k/internal/app/structs/models"
	"gorm.io/gorm"
)

func SeedTutorials(tx *gorm.DB) error {
	tutorials := []models.Tutorial{
		{
			Title:   "Cara Melakukan Pendaftaran",
			Content: "Buka halaman pendaftaran, pilih role Anda, dan isi formulir dengan data yang valid.",
		},
		{
			Title:   "Cara Memesan Produk UMKM",
			Content: "Masuk ke halaman produk, cari produk yang Anda inginkan, masukkan jumlahnya, lalu selesaikan pembayaran.",
		},
	}

	for _, tutorial := range tutorials {
		if err := tx.Create(&tutorial).Error; err != nil {
			return err
		}
	}
	return nil
}
