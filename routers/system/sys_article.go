package system

import "github.com/gin-gonic/gin"

type ArticleRouter struct {

}

func (s *ArticleRouter) InitArticleRouter() *gin.Engine{
	router := gin.New()

	systemApi := router.Group("api/system/article")
	{
		systemApi.POST("/article/listArticles")
		systemApi.POST("/article/getArticleDetails")
		systemApi.POST("/article/searchArticle")
		systemApi.POST("/article/saveArticle")
		systemApi.POST("/article/publishArticle")
		systemApi.POST("/article/updateArticle")
		systemApi.POST("/article/recoverArticle")
		systemApi.POST("/article/listArticleTags")
	}

	return router
}
