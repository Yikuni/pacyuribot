package test

import "github.com/gin-gonic/gin"

type CrawlerRouter struct {
}

func (s *CrawlerRouter) InitTestCrawlerRouter(rGroup *gin.RouterGroup) {
	r := rGroup.Group("crawler")
	r.POST("crawl", crawlerAPI.Crawl)
}
