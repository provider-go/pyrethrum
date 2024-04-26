package api

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/pkg/go-logger"
	service "github.com/provider-go/pyrethrum/service/manage"
)

// ManagePluginDownload 插件下载
func ManagePluginDownload(ctx *gin.Context) {
	pluginName := ctx.PostForm("pluginName")
	logger.Info("ManagePluginDownload req:", "pluginName", pluginName)
	service.ManagePluginDownload(pluginName)

	ctx.JSON(200, gin.H{
		"message": "Hello download!",
	})
}

// ManagePluginInstall 插件安装
func ManagePluginInstall(ctx *gin.Context) {
	pluginName := ctx.PostForm("pluginName")
	logger.Info("ManagePluginDownload req:", "pluginName", pluginName)
	service.ManagePluginInstall(pluginName)
}
