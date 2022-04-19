package router

import (
	"Cloudflare-Assistant/router/api"
	"github.com/gin-gonic/gin"
)

func SetAllRouter(r *gin.Engine) {
	// API
	api.SetAPIRouter(r.Group("/api"))
}
