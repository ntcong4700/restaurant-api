package main

import (
	"log"

	"restaurant-api/config"
	"restaurant-api/models"
)

func main() {
	db := config.ConnectDatabase()

	err := db.AutoMigrate(&models.MenuItem{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration completed")
}