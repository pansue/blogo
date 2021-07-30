package routers

import (
	"blogo/routers/blog"
	"blogo/routers/system"
)

type RouterGroup struct {
	System			system.RouterGroup
	Blog			blog.RouterGroup
}

var RouterGroupApp = new(RouterGroup)