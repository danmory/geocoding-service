package main

import (
	"log"

	"github.com/danmory/geocoding-service/auth-service/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	app.RunApp()
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error reading .env file")
	}
}