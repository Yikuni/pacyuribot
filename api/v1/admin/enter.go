package admin

import "pacyuribot/service"

type APIGroup struct {
	CrawlerAPI
}

var (
	crawlerService = service.ServiceGroupApp.AdminServiceGroup.CrawlerService
)
