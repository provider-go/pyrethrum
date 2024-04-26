package service

import "github.com/provider-go/pkg/go-logger"

// ManagePluginDownload 插件下载
func ManagePluginDownload(pluginName string) {

	logger.Info("ManagePluginDownload req:", "pluginName", pluginName)

}

// ManagePluginInstall 插件安装
func ManagePluginInstall(pluginName string) {

	logger.Info("ManagePluginDownload req:", "pluginName", pluginName)
}
