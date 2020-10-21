package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang-functions/enums"
	"golang-functions/models"
	"log"
	"net/http"
)

func AdminIndex(c *gin.Context) {

	adminName := sessions.Default(c).Get(enums.SessionAdminName)
	fmt.Println(adminName)

	m := &models.Feature{}
	features, count, err := m.Search(0, 100)
	if err != nil {
		log.Println(err)
	}

	c.HTML(http.StatusOK, "admin.index.html", gin.H{
		"features": features,
		"count": count,
	})
}

