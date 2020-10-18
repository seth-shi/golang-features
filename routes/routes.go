package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes() *gin.Engine {

	var router = gin.Default()

	webRoutes(router)
	adminRoutes(router)

	return router
}
