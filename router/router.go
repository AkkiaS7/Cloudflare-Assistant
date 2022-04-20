package router

import (
	"Cloudflare-Assistant/controller"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Server *gin.Engine
}

func InitRouter() *Server {
	server := &Server{}
	server.Server = gin.Default()

	server.RegisterRouter()

	return server
}

// Run 启动服务
func (s *Server) Run() {
	s.Server.Run(":8080")
}

// RegisterRouter 注册路由
func (s *Server) RegisterRouter() {
	controller.InitUserController(s.Server.Group("/user"))
	controller.InitOauth2Controller(s.Server.Group("/oauth2"))
	controller.InitApiController(s.Server.Group("/api"))
	controller.InitDashBoardController(s.Server.Group("/"))
}
