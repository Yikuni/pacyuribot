package test

import api "pacyuribot/api/v1"

type RouterGroup struct {
	CrawlerRouter
	PocketbaseRouter
}

var (
	crawlerAPI    = api.APIGroupApp.TestAPIGroup.CrawlerAPI
	pocketbaseAPI = api.APIGroupApp.TestAPIGroup.PocketBaseAPI
)
