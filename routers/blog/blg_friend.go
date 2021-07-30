package blog

import "github.com/gin-gonic/gin"

type FriendRouter struct {

}

func (f *FriendRouter) InitFriendRouter() *gin.Engine{
	router := gin.New()

	blogApi := router.Group("api/blog/friend")
	{
		blogApi.POST("/listFriend")
	}

	return router
}
