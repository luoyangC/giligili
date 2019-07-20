package main

import (
	"giligili/settings"
	"giligili/server"
)

func main() {
	// 从配置文件读取配置
	settings.Init()

	// 装载路由
	r := server.NewRouter()
	r.Run(":3000")
}
