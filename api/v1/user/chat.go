package user

import (
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"io"
	"pacyuribot/model/common/response"
)

type ChatAPI struct {
}

// Completions
// @Summary 流式对话，请求格式和chatgpt相同，但响应格式为纯文本
// @Router /user/chat/threads/runs
func (a *ChatAPI) Completions(c *gin.Context) {
	// 设置流式响应
	c.Header("Content-Type", "text/event-stream; charset=utf-8")
	c.Header("Transfer-Encoding", "chunked")

	var req openai.CreateThreadAndStreamRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.InvalidRequestFormat(c)
		return
	}
	c.Stream(func(w io.Writer) bool {
		assistantService.Chat(req.AssistantID, req.CreateThreadAndRunRequest.Thread.Messages,
			func(event openai.AssistantStreamEvent) {
				if len(event.Content) > 0 {
					w.Write([]byte(event.Content[0].Text.Value))
					ginW := w.(gin.ResponseWriter)
					ginW.Flush()
				}
			}, func() {
				//logger.Debug("finish")
				// TODO save the conversation
			})

		return false
	})
}
