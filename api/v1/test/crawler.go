package test

import (
	"github.com/gin-gonic/gin"
	"pacyuribot/core/crawler"
	"pacyuribot/logger"
	"pacyuribot/model/admin/request"
	"pacyuribot/model/common/response"
)

type CrawlerAPI struct {
}

func (a *CrawlerAPI) Crawl(c *gin.Context) {
	var config request.DefaultCrawlerConfig
	err := c.ShouldBindJSON(&config)
	if err != nil {
		response.InvalidRequestFormat(c)
		logger.Error(err.Error())
		return
	}
	crawler.GetTestCrawler(config).Run()
	response.OkWithMessage("正在执行", c)
}
