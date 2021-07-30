package system

import "github.com/gin-gonic/gin"

type ArticleRouter struct {

}

func (s *ArticleRouter) InitArticleRouter() *gin.Engine{
	router := gin.New()

	systemApi := router.Group("api/system/article")
	{
		systemApi.POST("/listArticles")
		systemApi.POST("/getArticleDetails")
		systemApi.POST("/searchArticle")
		systemApi.POST("/saveArticle")
		systemApi.POST("/publishArticle")
		systemApi.POST("/updateArticle")
		systemApi.POST("/recoverArticle")
		systemApi.POST("/listArticleTags")
	}

	return router
}
