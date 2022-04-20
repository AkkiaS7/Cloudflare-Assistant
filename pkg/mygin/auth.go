package mygin

import (
	"github.com/gin-gonic/gin"
	"strings"
)

type AuthOptions struct {
	Role     string
	Redirect string
}

func Authorize(opt AuthOptions) func(c *gin.Context) {
	return func(c *gin.Context) {
		token, _ := c.Cookie("token")
		if strings.TrimSpace(token) == "" {
			c.Redirect(302, opt.Redirect)
			c.Abort()
			return
		}
	}
}
