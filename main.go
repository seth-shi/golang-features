package main

import "golang-functions/routes"

func main() {

	router := routes.RegisterRoutes()

	router.Run("0.0.0.0:8888")
}