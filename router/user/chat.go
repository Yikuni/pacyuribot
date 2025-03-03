package user

import "github.com/gin-gonic/gin"

type ChatRouter struct {
}

func (c *ChatRouter) InitializeChatRouter(rGroup *gin.RouterGroup) {
	r := rGroup.Group("chat")
	r.POST("completion/:userID", chatAPI.Completions)
}
