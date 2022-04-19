package main

import "Cloudflare-Assistant/server"

func main() {
	go server.Start()
	select {}
}
