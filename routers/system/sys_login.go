package system

import "github.com/gin-gonic/gin"

type LoginRouter struct {

}

func (s *LoginRouter)InitLoginRouter() *gin.Engine {
	router := gin.New()

	systemApi := router.Group("api/system")
	{
		systemApi.POST("/auth/login")
	}

	return router
}