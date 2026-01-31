package database

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// Init MySQL connection
func newMysqlDB() (*gorm.DB, error) {
	_ = godotenv.Load() // optional, bisa dipindah ke main.go

	dbUser := os.Getenv("DBUSER")
	dbPassword := os.Getenv("DBPASSWORD")
	dbHost := os.Getenv("DBHOST")
	dbPort := os.Getenv("DBPORT")
	dbName := os.Getenv("DBNAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

// ✅ FUNCTION YANG KAMU MINTA
func GetDB() *gorm.DB {
	once.Do(func() {
		var err error
		db, err = newMysqlDB()
		if err != nil {
			panic("failed to connect database: " + err.Error())
		}
	})
	return db
}
