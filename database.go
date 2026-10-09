package main

import (
"log"
"gorm.io/driver/postgres"
"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"

db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
if err != nil {
log.Fatal("Database Connection Failed:", err)
}
DB = db

DB.AutoMigrate(&Student{})

log.Println("Database Connected Successfully")
}
