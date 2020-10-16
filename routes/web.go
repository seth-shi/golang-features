package routes

import (
	"github.com/gin-gonic/gin"
	. "golang-functions/controllers"
)

func webRoutes(router *gin.Engine)  {

	router.GET("/", Index)
	router.GET("ping", Ping)
}