package controller

import (
	"Cloudflare-Assistant/pkg/mygin"
	"github.com/gin-gonic/gin"
)

type userController struct {
	r gin.IRoutes
}

func InitUserController(r gin.IRoutes) {
	uc := &userController{r: r}
	uc.registerRouter()
}

func (uc *userController) registerRouter() {
	uc.r.Use(mygin.Authorize(mygin.AuthOptions{
		Role:     "user",
		Redirect: "/login",
	}))
	uc.r.GET("", uc.userHome)
}

func (uc *userController) userHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
