package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes() *gin.Engine {

	var router = gin.Default()

	webRoutes(router)

	return router
}
