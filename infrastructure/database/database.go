package database

import (
	"fatura-yonetim-sistemi/infrastructure/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

// Database instance
type Dbinstance struct {
	Db *gorm.DB
}

var Database Dbinstance

// ConnectDb function to connect to database
func ConnectDb() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.ConfigEnv("DB_HOST"),
		config.ConfigEnv("DB_USER"),
		config.ConfigEnv("DB_PASSWORD"),
		config.ConfigEnv("DB_NAME"),
		config.ConfigEnv("DB_PORT"),
	)
	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	err = db.AutoMigrate()
	if err != nil {
		log.Fatal("Failed to migrate. \n", err)
		os.Exit(2)
	}
	Database = Dbinstance{
		Db: db,
	}
}
