package user

import "pacyuribot/api/v1"

type RouterGroup struct {
	ChatRouter
}

var (
	chatAPI = api.APIGroupApp.UserAPIGroup.ChatAPI
)
