package main

import (
	"Cloudflare-Assistant/router"
	"Cloudflare-Assistant/service"
)

var (
	server *router.Server
)

func init() {
	service.InitService()
	server = router.InitRouter()
}

func main() {
	server.Run()
	select {}
}
