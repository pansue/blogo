package blog

import "github.com/gin-gonic/gin"

type FriendRouter struct {

}

func (f *FriendRouter) InitFriendRouter(Router *gin.RouterGroup) {
	blogApi := Router.Group("api/blog/friend")
	{
		blogApi.POST("/listFriend")
	}
}
