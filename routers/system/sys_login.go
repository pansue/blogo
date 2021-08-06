package system

import (
	"blogo/api"

	"github.com/gin-gonic/gin"
)

type LoginRouter struct {

}

func (s *LoginRouter)InitLoginRouter(Router *gin.RouterGroup) {
	systemApi := Router.Group("api/system/auth")
	{
		systemApi.POST("/login", api.ApiGroupApp.SystemApiGroup.LoginApi.Login)
	}
}
