package initialize

import (
	"gin/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router:=gin.Default()

	systemRouter:=router.RouterGroupApp.System

	PrivateGroup:=Router.Group("")

	{
		systemRouter.InitApiRouter(PrivateGroup)		//注册功能api路由
		systemRouter.InitLoginApiRouter(PrivateGroup)	//注册登录api路由
		systemRouter.InitUploadRouter(PrivateGroup)	//注册上传api路由
		systemRouter.InitUserRouter(PrivateGroup)	//注册系统用户api路由
	}

	return Router
}
