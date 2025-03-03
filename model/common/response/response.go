package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pacyuribot/logger"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 1
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	if code == ERROR {
		c.AbortWithStatusJSON(http.StatusInternalServerError, Response{
			Code: code,
			Data: data,
			Msg:  msg,
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "出错啦", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
	logger.Error(message)
}

func NoAuth(c *gin.Context) {
	c.JSON(http.StatusNonAuthoritativeInfo, Response{
		ERROR, nil, "Unauthorized",
	})
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func InvalidRequestFormat(c *gin.Context) {
	Result(ERROR, nil, "Invalid request format", c)
}
