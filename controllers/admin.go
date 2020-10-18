package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminLogin(c *gin.Context) {

	c.HTML(http.StatusOK, "admin.login.html", gin.H{})
}

func AdminIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.index.html", gin.H{})
}