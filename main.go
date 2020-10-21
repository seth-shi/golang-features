package main

import (
	"github.com/joho/godotenv"
	"golang-functions/models"
	"golang-functions/routes"
	"os"
)

func init() {

	var err error

	err = godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	models.InitEs()
}

func main() {

	router := routes.RegisterRoutes()

	router.Run("0.0.0.0:" + os.Getenv("SERVE_PORT"))
}
