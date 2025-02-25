package admin

import api "pacyuribot/api/v1"

type RouterGroup struct {
	CrawlerRouter
	DatasourceRouter
}

var (
	crawlerAPI    = api.APIGroupApp.AdminAPIGroup.CrawlerAPI
	datasourceAPI = api.APIGroupApp.AdminAPIGroup.DatasourceAPI
)
