package admin

import (
	"pacyuribot/service"
	"pacyuribot/service/assistant"
)

type APIGroup struct {
	CrawlerAPI
}

var (
	crawlerService                               = service.ServiceGroupApp.AdminServiceGroup.CrawlerService
	assistantService  assistant.AssistantService = &(service.ServiceGroupApp.AssistantServiceGroup.ChatGPTService)
	datasourceService                            = service.ServiceGroupApp.AdminServiceGroup.DatasourceService
)
