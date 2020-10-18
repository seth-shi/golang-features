package routes

import (
	"github.com/gin-gonic/gin"
	. "golang-functions/controllers"
	"golang-functions/middlewares"
)

func adminRoutes(router *gin.Engine)  {

	router.GET("/admin/login", AdminLogin)
	router.POST("/admin/login", HandleAdminLogin)


	adminRouter := router.Group("/admin")
	adminRouter.Use(middlewares.AdminAuth)
	{
		adminRouter.GET("", AdminIndex)
	}
}