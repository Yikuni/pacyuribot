package service

import (
	"pacyuribot/service/admin"
	"pacyuribot/service/assistant"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	AdminServiceGroup     admin.AdminServiceGroup
	AssistantServiceGroup assistant.AssistantServiceGroup
}
