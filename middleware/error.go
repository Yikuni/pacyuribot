package middleware

import (
	"github.com/gin-gonic/gin"
	"pacyuribot/model/common/response"
)

func DefaultErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // 先执行后续中间件和处理函数

		// 检查是否有错误
		if len(c.Errors) > 0 {
			// 统一处理错误
			response.FailWithMessage(c.Errors.Last().Error(), c)
		}
	}
}
