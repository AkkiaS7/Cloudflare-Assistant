package api

import "github.com/gin-gonic/gin"

func SetAPIRouter(r *gin.RouterGroup) {
	SetZoneRouter(r.Group("/zones"))
}
