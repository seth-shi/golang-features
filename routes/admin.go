package routes

import (
	"github.com/gin-gonic/gin"
	. "golang-functions/controllers"
)

func adminRoutes(router *gin.Engine)  {

	router.GET("/admin/login", AdminLogin)
	router.GET("/admin", AdminIndex)
}