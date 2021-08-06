package system

import (
	"blogo/models/request"
	"blogo/services/account_service"
	"blogo/utils"
	"github.com/gin-gonic/gin"
)

type LoginApi struct {
}

func (b *LoginApi) Login(c *gin.Context) {
	var loginInfo request.Login
	c.ShouldBindJSON(&loginInfo)
	if loginInfo.Username == "" || loginInfo.Password == "" {
		c.JSON(401, gin.H{
			"登录失败": "账号或密码不能为空",
		})
		return
	}
	_, err := account_service.CheckAuth(loginInfo.Username, loginInfo.Password)
	if err != nil {
		c.JSON(401, gin.H{
			"登录失败": "账号或密码错误",
		})
		return
	}
	accessToken, expireTime, err := utils.GetToken(123, loginInfo.Username)
	if err != nil {
		c.JSON(500, gin.H{
			"系统错误": "获取Token失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"accessToken": accessToken,
		"expiryTime":  expireTime,
		"user":        loginInfo.Username,
	})
}
