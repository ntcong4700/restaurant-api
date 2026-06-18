package main

import (
	"log"

	"github.com/joho/godotenv"

	"restaurant-api/config"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found")
	}

	config.ConnectDatabase()

	log.Println("App started")
}
