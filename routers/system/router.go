package system

import "github.com/gin-gonic/gin"

// InitSystemRouter initialize system routing
func InitSystemRouter() *gin.Engine{
	router := gin.New()

	systemApi := router.Group("api/system")
	{
		//public api
		systemApi.POST("/auth/login")

		//private api
		systemApi.POST("/article/listArticles")
		systemApi.POST("/article/getArticleDetails")
		systemApi.POST("/article/searchArticle")
		systemApi.POST("/article/saveArticle")
		systemApi.POST("/article/publishArticle")
		systemApi.POST("/article/updateArticle")
		systemApi.POST("/article/recoverArticle")
		systemApi.POST("/article/listArticleTags")

		systemApi.POST("/friend/listFriends")
		systemApi.POST("/friend/addFriends")
		systemApi.POST("/friend/deleteFriend")
	}

	return router
}