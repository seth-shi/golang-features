package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang-functions/utils"
	"net/http"
	"os"
)

func AdminLogin(c *gin.Context) {

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

	session.Set("admin_name", os.Getenv("ADMIN_USERNAME"))
	c.Redirect(http.StatusFound, "/admin")
}

func AdminIndex(c *gin.Context) {

	fmt.Println(sessions.Default(c).Get("admin_name"))
	c.HTML(http.StatusOK, "admin.index.html", gin.H{})
}
