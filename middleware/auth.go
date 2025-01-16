package middleware

import (
	"github.com/gin-gonic/gin"
	"pacyuribot/model/common/response"
	utils "pacyuribot/utils/pocketbase"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := utils.Auth(c.GetHeader("Authorization"))
		if err != nil {
			response.NoAuth(c)
			c.Abort()
			return
		}
		c.Set("id", id)
	}
}
