package main

import (
	"github.com/joho/godotenv"
	"github.com/jrschmidtt/api-gin/database"
	"github.com/jrschmidtt/api-gin/routes"
	"log"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file please create one with the same structure as .env.sample")
	}
	database.ConnectPostgresDB()
	routes.HandleRequest()
}