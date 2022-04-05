package main

import (
	"log"

	router "github.com/t-east/pr-sounds-backend/src/drivers/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env not found")
	}

	log.Println("Server running...")
	router.Serve()
}
