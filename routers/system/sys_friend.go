package system

import "github.com/gin-gonic/gin"

type FriendRouter struct {

}

func (s *FriendRouter)InitFriendRouter() *gin.Engine {
	router := gin.New()

	systemApi := router.Group("api/system/friend")
	{
		systemApi.POST("/friend/listFriends")
		systemApi.POST("/friend/addFriends")
		systemApi.POST("/friend/deleteFriend")
	}

	return router
}