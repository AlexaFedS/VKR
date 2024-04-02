package main

import (
	"log"

	"Lab1/internal/api"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	log.Println("Application start!")
	api.StartServer()
	dsn := "host=localhost user=postgres password=test dbname=test port=5432 sslmode=disable"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	log.Println("Status: OK!")
	log.Println("Application terminated!")
}
