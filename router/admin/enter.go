package admin

import api "pacyuribot/api/v1"

type RouterGroup struct {
	CrawlerRouter
}

var (
	crawlerAPI = api.APIGroupApp.AdminAPIGroup.CrawlerAPI
)
