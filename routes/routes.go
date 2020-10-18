package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {

	var router = gin.Default()
	startSession(router)
	loadStaticResources(router)

	webRoutes(router)
	adminRoutes(router)

	return router
}

func startSession(router *gin.Engine)  {

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("sessionid", store))
}

func loadStaticResources(router *gin.Engine)  {

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("views/*")
}
