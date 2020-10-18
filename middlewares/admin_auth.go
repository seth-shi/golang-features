package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminAuth(c *gin.Context) {

	session := sessions.Default(c)

	if session.Get("admin_name") == nil {
		session.AddFlash("请先登录后台账号")
		_ = session.Save()
		c.Redirect(http.StatusFound, "/admin/login")
	}
}
