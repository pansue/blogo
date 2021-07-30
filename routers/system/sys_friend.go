package system

import "github.com/gin-gonic/gin"

type FriendRouter struct {

}

func (s *FriendRouter)InitFriendRouter(Router *gin.RouterGroup) {
	systemApi := Router.Group("api/system/friend")
	{
		systemApi.POST("/listFriends")
		systemApi.POST("/addFriends")
		systemApi.POST("/deleteFriend")
	}
}
