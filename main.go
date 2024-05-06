package main

import (
	"github.com/provider-go/pyrethrum/config"
	"github.com/provider-go/pyrethrum/core/global"
	"github.com/provider-go/pyrethrum/router"
	"os"
)

func main() {

	// 加载配置
	config.Start()
	// 初始化
	global.Initialize()
	global.RootDir, _ = os.Getwd()

	// 启动服务并注册
	router.RouterStart()
}
