package utils

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Flashes(c *gin.Context,vars... string) (string, bool) {

	session := sessions.Default(c)

	var flashes []interface{}
	if len(vars) > 0 {
		flashes = session.Flashes(vars[0])
	} else {
		flashes = session.Flashes()
	}


	if len(flashes) == 0 {
		return "", false
	}

	msg := ""
	// 取完所有 flashes
	for _, item := range flashes {

		if m, ok := item.(string); ok {
			msg += m
		}
	}
	_ = session.Save()

	return msg, true
}