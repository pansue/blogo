package system

import "github.com/gin-gonic/gin"

type LoginRouter struct {

}

func (s *LoginRouter)InitLoginRouter(Router *gin.RouterGroup) {
	systemApi := Router.Group("api/system/login")
	{
		systemApi.POST("/auth/login")
	}
}