package controller

import "github.com/gin-gonic/gin"

type apiController struct {
	r gin.IRoutes
}

func InitApiController(r gin.IRoutes) {
	uc := &apiController{r: r}
	uc.registerRouter()
}

func (uc *apiController) registerRouter() {
	uc.r.Use()
}

func (uc *apiController) userHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
