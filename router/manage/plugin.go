package manageRouter

import (
	"github.com/gin-gonic/gin"
	api "github.com/provider-go/pyrethrum/api/manage"
)

func InitPluginRouter(Router *gin.RouterGroup) gin.IRouter {
	pluginRouter := Router.Group("plugin")
	{
		pluginRouter.Any("/download", api.ManagePluginDownload)
		pluginRouter.POST("/install", api.ManagePluginInstall)
	}

	return pluginRouter
}
