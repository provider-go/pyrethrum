package router

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/content"
	"github.com/provider-go/pyrethrum/core/global"
	"github.com/provider-go/pyrethrum/core/protocol"
)

func ImportPlugin(r *gin.Engine) {
	publicGroup := r.Group("")
	PluginInit(publicGroup, content.CreatePluginAndDB(global.DB))

}

func PluginInit(group *gin.RouterGroup, Plugin ...protocol.Plugin) {
	for i := range Plugin {
		PluginGroup := group.Group(Plugin[i].RouterPath())
		Plugin[i].Register(PluginGroup)
	}
}
