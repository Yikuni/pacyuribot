package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"io"
	"pacyuribot/model/common/response"
)

type ChatAPI struct {
}

// Completions
// @Summary 流式对话，接口和chatgpt几乎相同
// @Router /user/chat/completions/:userID
func (a *ChatAPI) Completions(c *gin.Context) {
	// 设置流式响应
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Header("Transfer-Encoding", "chunked")

	var req openai.CreateThreadAndStreamRequest
	err := c.ShouldBindJSON(req)
	if err != nil {
		response.InvalidRequestFormat(c)
		return
	}
	assistantService.Chat(req.AssistantID, req.CreateThreadAndRunRequest.Thread.Messages,
		func(event openai.AssistantStreamEvent) {
			jsonBytes, err := json.Marshal(event)
			if err != nil {
				response.InvalidRequestFormat(c)
				return
			}
			c.Stream(func(w io.Writer) bool {
				w.Write(jsonBytes)
				w.Write([]byte("\n\n"))
				return true
			})
		}, func() {
			response.Ok(c)
		})
}
