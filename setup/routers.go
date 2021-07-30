package setup

import (
	"blogo/routers"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine{
	Router := gin.Default()

	systemRouter := routers.RouterGroupApp.System
	blogRouter := routers.RouterGroupApp.Blog

	// public router
	PublicGroup := Router.Group("")
	{
		systemRouter.InitLoginRouter(PublicGroup)
	}

	// private router
	PrivateGroup := Router.Group("")
	{
		systemRouter.InitArticleRouter(PrivateGroup)
		systemRouter.InitFriendRouter(PrivateGroup)

		blogRouter.InitArticleRouter(PrivateGroup)
		blogRouter.InitFriendRouter(PrivateGroup)
	}

	return Router
}
