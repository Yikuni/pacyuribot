package utils

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

func Ask() {
	apiKey := "sk-DBiCHH4YHkeFuctIaRGlpwOs6uYU40Hh252cvH44gxk67FFg"
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = "https://api.moonshot.cn/v1"
	client := openai.NewClientWithConfig(config)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: "moonshot-v1-8k",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "你好",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)

}

func UploadFile() {
	apiKey := "sk-rW36IUx32nZrNdKPIjl71xPz6P583UyZVeCOB9Qx76Aph8cR"
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = "https://api.moonshot.cn/v1"
	client := openai.NewClientWithConfig(config)

	_, err := client.CreateFile(context.Background(), openai.FileRequest{
		FileName: "spigot-reflect文档",
		FilePath: "README.md",
		Purpose:  "file-extract",
	})
	if err != nil {
		return
	}
}
