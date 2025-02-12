package utils

import (
	"github.com/sashabaranov/go-openai"
	"net/http"
	"net/url"
	"pacyuribot/global"
)

func GetChatGPTClient() *openai.Client {
	config := openai.DefaultConfig(global.Config.Chatgpt.Key)
	if global.Config.Chatgpt.Proxy != "" {
		proxyUrl, err := url.Parse(global.Config.Chatgpt.Proxy)
		if err != nil {
			panic(err)
		}
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}
		config.HTTPClient = &http.Client{
			Transport: transport,
		}
	}

	return openai.NewClientWithConfig(config)
}
