package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/provider-go/pkg/go-logger"
	util "github.com/provider-go/pkg/util/network"
	manageRouter "github.com/provider-go/pyrethrum/router/manage"
	"github.com/spf13/viper"
)

// RouterStart 启动路由
func RouterStart() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	ImportPlugin(r)

	ManageGroup := r.Group("manage")
	{
		manageRouter.InitPluginRouter(ManageGroup)
	}

	printRoutes(r)

	logger.Info("API server start", "listen", util.GetHostIp()+viper.GetString("service.listen"))
	err := r.Run(viper.GetString("service.listen"))
	if err != nil {
		return
	}
}

func printRoutes(engine *gin.Engine) {
	routes := engine.Routes()
	for _, route := range routes {
		fmt.Printf("%s %s -> %s\n", route.Method, route.Path, route.Handler)
	}
}
