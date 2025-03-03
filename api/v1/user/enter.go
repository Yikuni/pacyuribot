package user

import (
	"pacyuribot/service"
	"pacyuribot/service/assistant"
)

type APIGroup struct {
	ChatAPI
}

var (
	assistantService assistant.AssistantService = &(service.ServiceGroupApp.AssistantServiceGroup.ChatGPTService)
)
