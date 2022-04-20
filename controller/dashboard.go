package controller

import "github.com/gin-gonic/gin"

type DashboardController struct {
	r gin.IRoutes
}

func InitDashBoardController(r *gin.RouterGroup) {
	dc := DashboardController{r: r}
	dc.registerRouter()
}

func (dc *DashboardController) registerRouter() {
	dc.r.GET("/", dc.index)
	dc.r.GET("/login", dc.login)
}

// index 主页
func (dc *DashboardController) index(c *gin.Context) {
	c.HTML(200, "dashboard/index.html", gin.H{
		"title": "Dashboard",
	})
}

// login 登入页面
func (dc *DashboardController) login(c *gin.Context) {
	c.HTML(200, "dashboard/login.html", gin.H{
		"title": "Login",
	})
}
