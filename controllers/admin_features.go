package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang-functions/models"
	"golang-functions/requests"
	"log"
	"net/http"
)

func AdminFeatureEdit(c *gin.Context) {

	id := c.Param("id")

	m, err := models.Feature{Id: id}.Find()
	if err != nil {
		log.Println(err)
	}

	c.HTML(http.StatusOK, "admin.feature.edit.html", gin.H{
		"feature": m,
	})
}

func AdminFeatureUpdate(c *gin.Context) {

	id := c.Param("id")

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

	m, err := models.Feature{Id: id}.Find()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg": "无效的模型",
		})
		return
	}

	m.Title = form.Title
	m.Description = form.Description
	m.Code = form.Code
	// TODO 更新失败
	err = models.Update(m)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": "修改成功",
	})
}
