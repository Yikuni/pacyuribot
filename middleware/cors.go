package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers",
			"Content-Type,AccessToken,X-CSRF-Token,Openai-Beta,Authorization,X-Stainless-Runtime,"+
				"X-Stainless-Timeout,X-Stainless-Runtime-Version,X-Stainless-arch,X-Stainless-Lang,"+
				"X-Stainless-Os,X-Stainless-Package-Version,X-Stainless-Retry-Count")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
