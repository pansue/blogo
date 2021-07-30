package system

import "github.com/gin-gonic/gin"

type ArticleRouter struct {

}

func (s *ArticleRouter) InitArticleRouter(Router *gin.RouterGroup){
	systemApi := Router.Group("api/system/article")
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
}
