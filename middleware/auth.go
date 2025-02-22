package middleware

import (
	"github.com/gin-gonic/gin"
	"pacyuribot/model/common/response"
	"pacyuribot/utils"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := utils.Auth(c.GetHeader("Authorization"))
		if err != nil {
			response.NoAuth(c)
			c.Abort()
			return
		}
		c.Set("userID", id)
	}
}
