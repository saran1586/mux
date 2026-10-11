package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDB() {

	// Connection string comes from DATABASE_URL; the fallback is a local development database.
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database Connection Failed:", err)
	}
	DB = db

	DB.AutoMigrate(&Student{})

	log.Println("Database Connected Successfully")
}
