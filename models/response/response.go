package response

import "github.com/gin-gonic/gin"

// GinContextE 是对Gin.Context 的扩展
type GinContextE struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func (g GinContextE) Response(httpCode, errCode int, msg string, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
}

func (g GinContextE) Ok() {
	g.Response(200, SUCCESS, "操作成功", map[string]interface{}{})
}

func (g GinContextE) OkWithMessage(message string) {
	g.Response(200, SUCCESS, message, map[string]interface{}{})
}

func (g GinContextE) OkWithData(data interface{}) {
	g.Response(200, SUCCESS, "操作成功", data)
}

func (g GinContextE) OkWithDetailed(data interface{}, message string) {
	g.Response(200, SUCCESS, message, data)
}

func (g GinContextE) Fail(httpCode int) {
	g.Response(httpCode, ERROR, "操作失败", map[string]interface{}{})
}

func (g GinContextE) FailWithMessage(httpCode int, message string) {
	g.Response(httpCode, ERROR, message, map[string]interface{}{})
}

func (g GinContextE) FailWithData(httpCode int, data interface{}) {
	g.Response(httpCode, ERROR, "操作失败", data)
}

func (g GinContextE) FailWithDetailed(httpCode int, data interface{}, message string) {
	g.Response(httpCode, ERROR, message, data)
}
