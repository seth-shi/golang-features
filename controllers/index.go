package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang-functions/models"
	"golang-functions/requests"
	"net/http"
)

func Index(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{})
}


func FeaturesCreate(c *gin.Context) {


	c.HTML(http.StatusOK, "features.create.html", gin.H{

	})
}

func FeatureStore(c *gin.Context) {

	var form requests.FeatureRequest
	if err := c.ShouldBind(&form); err != nil {

		msg := err.Error()
		if v, ok := err.(validator.ValidationErrors); ok {
			msg = form.GetError(v)
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg": msg,
		})
		return
	}

	m := models.Feature{}
	m.Title = form.Title
	m.Description = form.Description
	m.Code = form.Code
	m.ViewCount = 0
	err := models.Create(m)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": "提交成功,等待审核",
	})
}
