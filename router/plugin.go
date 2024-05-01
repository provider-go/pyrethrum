package router

import (
	"github.com/gin-gonic/gin"
	attachfile "github.com/provider-go/attach-file"
	"github.com/provider-go/content"
	"github.com/provider-go/district"
	"github.com/provider-go/pyrethrum/core/global"
	"github.com/provider-go/pyrethrum/core/protocol"
	"github.com/provider-go/user"
)

func ImportPlugin(r *gin.Engine) {
	publicGroup := r.Group("")
	PluginInit(publicGroup, content.CreatePluginAndDB(global.DB),
		user.CreatePluginAndDB(global.DB),
		district.CreatePluginAndDB(global.DB),
		attachfile.CreatePluginAndDB(global.DB))

}

func PluginInit(group *gin.RouterGroup, Plugin ...protocol.Plugin) {
	for i := range Plugin {
		PluginGroup := group.Group(Plugin[i].RouterPath())
		Plugin[i].Register(PluginGroup)
	}
}
