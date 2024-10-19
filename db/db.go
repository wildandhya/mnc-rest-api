package db

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/wildandhya/mnc-rest-api/internal/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	errENV := godotenv.Load()
	if errENV != nil {
		panic(errENV)
	}

	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	createDatabaseIfNotExist(db, DB_NAME)

	db.AutoMigrate(&entity.User{}, &entity.TopUp{}, &entity.Payment{}, &entity.Transfer{})

	connection, err := db.DB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	connection.SetMaxIdleConns(10)
	connection.SetMaxOpenConns(100)
	connection.SetConnMaxLifetime(time.Second * time.Duration(300))

	return db

}

func createDatabaseIfNotExist(db *gorm.DB, dbName string) {
	createDBQuery := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)
	if err := db.Exec(createDBQuery).Error; err != nil {
		log.Fatalf("Failed to create database %s: %v", dbName, err)
	}
	fmt.Printf("Database %s checked/created successfully.\n", dbName)
}
