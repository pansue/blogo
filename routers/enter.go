package routers

import (
	"blogo/routers/blog"
	"blogo/routers/system"
)

type Server struct {
	System			system.RouterGroup
	Blog			blog.RouterGroup
}