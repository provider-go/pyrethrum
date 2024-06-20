package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/provider-go/pkg/logger"
	util "github.com/provider-go/pkg/util"
	"github.com/spf13/viper"
)

// StartRouter 启动路由
func StartRouter() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	ImportPlugin(r)

	printRoutes(r)

	logger.Info("API server start", "listen", util.GetHostIp()+viper.GetString("service.listen"))
	err := r.Run(viper.GetString("service.listen"))
	if err != nil {
		logger.Error("API server start", "err", err)
		return
	}
}

func printRoutes(engine *gin.Engine) {
	routes := engine.Routes()
	for _, route := range routes {
		fmt.Printf("%s %s -> %s\n", route.Method, route.Path, route.Handler)
	}
}
