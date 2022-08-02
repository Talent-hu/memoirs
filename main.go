package main

import (
	"fmt"
	_ "memoirs/config"
	_ "memoirs/docs"
	"memoirs/global"
	"memoirs/pkg/cache"
	"memoirs/router"
)

// @title 后台系统
// @version 1.0.0
// @description gin开发的后台管理系统
func main() {
	cache.InitCache()
	addr := fmt.Sprintf(":%s", global.Config.System.Port)
	server := router.InitRouter()
	_ = server.Run(addr)
}
