package server

import (
	"Cloudflare-Assistant/handler"
	"Cloudflare-Assistant/router"
	"github.com/gin-gonic/gin"
)

var (
	r *gin.Engine
)

func init() {
	r = gin.Default()
	handler.Init()
}
func Start() {
	router.SetAllRouter(r)
	r.Run() // listen and serve on
}
