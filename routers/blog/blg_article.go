package blog

import "github.com/gin-gonic/gin"

type ArticleRouter struct {

}

// InitArticleRouter initialize blog routing
func (s *ArticleRouter)InitArticleRouter() *gin.Engine {
	router := gin.New()

	blogApi := router.Group("api/blog/article")
	{
		blogApi.POST("/listArticles")
		blogApi.POST("/ArticleDetails")
		blogApi.POST("/searchArticle")
		blogApi.POST("/listArticleTags")
	}

	return router
}