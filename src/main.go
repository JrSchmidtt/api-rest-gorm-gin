package main

import (
	"github.com/jrschmidtt/api-gin/database"
	"github.com/jrschmidtt/api-gin/routes"
)

func main() {
	database.ConnectPostgresDB()
	routes.HandleRequest()
}