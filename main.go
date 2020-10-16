package main

import (
	"github.com/joho/godotenv"
	"golang-functions/routes"
	"log"
	"os"
)

func init()  {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	router := routes.RegisterRoutes()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("views/*")

	router.Run("0.0.0.0:" + os.Getenv("SERVE_PORT"))
}