package main

import (
	"lcb-go/conf"
	"lcb-go/server"
)

func main() {
	// 加载配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()
	r.Run(":8080")
}
