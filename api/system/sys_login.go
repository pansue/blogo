package system

import (
	"blogo/models/request"
	"blogo/models/response"
	"blogo/services/account_service"
	"blogo/utils"
	"github.com/gin-gonic/gin"
)

type LoginApi struct {
}

func (b *LoginApi) Login(c response.GinContextE) {
	var loginInfo request.LoginRequest
	c.C.ShouldBindJSON(&loginInfo)
	if loginInfo.Username == "" || loginInfo.Password == "" {
		c.FailWithDetailed(401, gin.H{"登录失败": "账号或密码不能为空"}, "账号或密码不能为空")
		return
	}
	_, err := account_service.CheckAuth(loginInfo.Username, loginInfo.Password)
	if err != nil {
		c.FailWithDetailed(401, gin.H{"登录失败": "账号或密码错误"}, "账号或密码错误")
		return
	}
	accessToken, expireTime, err := utils.GetToken(123, loginInfo.Username)
	if err != nil {
		c.FailWithDetailed(200, gin.H{"系统错误": "获取Token失败"}, "获取Token失败")
		return
	}
	c.OkWithDetailed(response.LoginResponse{
		AccessToken: accessToken,
		ExpireTime:  expireTime,
		User:        loginInfo.Username},
		"登录成功")
}
