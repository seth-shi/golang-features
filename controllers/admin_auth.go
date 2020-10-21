package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang-functions/enums"
	"golang-functions/utils"
	"net/http"
	"os"
)

func AdminLogin(c *gin.Context) {

	session := sessions.Default(c)
	if session.Get(enums.SessionAdminName) != nil {
		c.Redirect(http.StatusFound, "/admin")
		return
	}

	h := gin.H{}
	h["_flash"], h["has_flash"] = utils.Flashes(c)
	h["username"], _ = utils.Flashes(c, "username")

	c.HTML(http.StatusOK, "admin.login.html", h)
}

func HandleAdminLogin(c *gin.Context) {

	session := sessions.Default(c)
	defer session.Save()

	_ = c.Request.ParseForm()
	username, password := c.Request.PostFormValue("username"), c.Request.PostFormValue("password")

	if len(username) == 0 || len(password) == 0 {
		session.AddFlash("用户名和密码不能为空")
		session.AddFlash(username, "username")
		c.Redirect(http.StatusFound, "/admin/login")
		return
	}
	if username != os.Getenv("ADMIN_USERNAME") ||
		password != os.Getenv("ADMIN_PASSWORD") {
		session.AddFlash("用户名或者密码错误")
		session.AddFlash(username, "username")

		c.Redirect(http.StatusFound, "/admin/login")
		return
	}

	session.Set(enums.SessionAdminName, os.Getenv("ADMIN_USERNAME"))
	c.Redirect(http.StatusFound, "/admin")
}