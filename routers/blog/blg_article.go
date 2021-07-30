package blog

import "github.com/gin-gonic/gin"

// InitBlogRouter initialize blog routing
func InitBlogRouter() *gin.Engine {
	router := gin.New()

	blogApi := router.Group("api/blog")
	{
		blogApi.POST("article/listArticles")
		blogApi.POST("article/getArticleDetails")
		blogApi.POST("article/searchArticle")
		blogApi.POST("article/listArticleTags")

		blogApi.POST("friend/listFriends")
	}

	return router
}