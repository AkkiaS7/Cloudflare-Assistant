package controller

import (
	"Cloudflare-Assistant/service"
	"github.com/gin-gonic/gin"
)

type oauth2Controller struct {
	r gin.IRoutes
}

func InitOauth2Controller(r gin.IRoutes) {
	oc := &oauth2Controller{r: r}
	oc.registerRouter()
}

func (oc *oauth2Controller) registerRouter() {
	r := oc.r
	flag := false
	if service.Conf.Oauth2.Github.Enable {
		flag = true
		r.GET("/github/callback", oc.githubCallback)
		r.GET("/github/login", oc.githubLogin)
	}
	if !flag {
		panic("no oauth2 enabled")
	}
}

func (oc *oauth2Controller) githubLogin(c *gin.Context) {
	c.HTML(200, "oauth2/login.html", nil)
}

func (oc *oauth2Controller) githubCallback(c *gin.Context) {
	c.HTML(200, "oauth2/callback.html", nil)
}
