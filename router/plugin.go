package router

import (
	"github.com/gin-gonic/gin"
	attachfile "github.com/provider-go/attach-file"
	"github.com/provider-go/manager"
	"github.com/provider-go/pyrethrum/core/global"
	"github.com/provider-go/pyrethrum/core/protocol"
	"github.com/provider-go/sms"
	"github.com/provider-go/sso"
	"github.com/spf13/viper"
)

func ImportPlugin(r *gin.Engine) {
	publicGroup := r.Group(viper.GetString("service.pre"))
	PluginInit(publicGroup,
		attachfile.CreatePluginAndDB(global.PluginConfig),
		sms.CreatePluginAndDB(global.PluginConfig),
		manager.CreatePluginAndDB(global.PluginConfig),
		sso.CreatePluginAndDB(global.PluginConfig),
	)
}

func PluginInit(group *gin.RouterGroup, Plugin ...protocol.Plugin) {
	for i := range Plugin {
		PluginGroup := group.Group(Plugin[i].RouterPath())
		Plugin[i].Register(PluginGroup)
	}
}
