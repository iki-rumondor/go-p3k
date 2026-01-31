package config

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

func NewMysqlDB() (*gorm.DB, error) {

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	dbUser := os.Getenv("DBUSER")
	dbPassword := os.Getenv("DBPASSWORD")
	dbHost := os.Getenv("DBHOST")
	dbPort := os.Getenv("DBPORT")
	dbName := os.Getenv("DBNAME")

	initialDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local&timeout=5s", dbUser, dbPassword, dbHost, dbPort)
	initialDB, err := gorm.Open(mysql.Open(initialDSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	var exists bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = '%s')", dbName)
	initialDB.Raw(query).Scan(&exists)

	if !exists {
		if err := initialDB.Exec("CREATE DATABASE " + dbName).Error; err != nil {
			return nil, err
		}
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=5s", dbUser, dbPassword, dbHost, dbPort, dbName)
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return gormDB, nil
}

func GetDB() *gorm.DB {
	once.Do(func() {
		var err error
		db, err = NewMysqlDB()
		if err != nil {
			panic(err)
		}
	})
	return db
}
