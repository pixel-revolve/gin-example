package router

import (
	"gin/router/system"
)

type RouterGroup struct {
	System system.RouterGroup

}

var RouterGroupApp =new(RouterGroup)