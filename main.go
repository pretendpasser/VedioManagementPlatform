package main

import (
	"VMP/conf"
	"VMP/server"
)

func main() {
	conf.Init()

	// 装载路由
	r := server.NewRouter()
	r.Run(":3000")
}