package admin

import "github.com/gin-gonic/gin"

type CrawlerRouter struct {
}

func (c *CrawlerRouter) InitCrawlerRouter(rGroup *gin.RouterGroup) {
	r := rGroup.Group("crawler")
	r.POST("crawl", crawlerAPI.Crawl)
}
