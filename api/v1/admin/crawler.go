package admin

import (
	"github.com/gin-gonic/gin"
	"pacyuribot/core/crawler"
	"pacyuribot/logger"
	"pacyuribot/model/admin/request"
	"pacyuribot/model/common/response"
)

type CrawlerAPI struct {
}

// Crawl
// @Summary		创建爬虫任务并执行
// @Security 	Auth
// @Param 		datasource datasourceID
// @Success 	200 {string} string "{"code":0,"data":{},"msg":"taskID"}"
// @Router 		/admin/crawl/crawl
func (a *CrawlerAPI) Crawl(c *gin.Context) {
	var config request.DefaultCrawlerConfig
	datasourceID := c.Query("datasource")
	userID := c.GetString("userID")
	err := c.ShouldBindJSON(&config)
	if err != nil {
		response.InvalidRequestFormat(c)
		logger.Error(err.Error())
		return
	}

	if datasourceID == "" {
		response.FailWithMessage("datasource cannot be nil", c)
		return
	}

	// create task
	taskID, err := crawlerService.CreateCrawlTask(userID, datasourceID, config)
	if err != nil {
		response.FailWithMessage("创建任务失败", c)
		logger.Error(err.Error())
		return
	}
	go func() {
		crawler.GetDefaultCrawler(config, userID, datasourceID).Run()
		err := crawlerService.SetCrawlTaskStatus(taskID)
		if err != nil {
			logger.Error("Failed to set crawl task status to completed: %s", err.Error())
			return
		}
		logger.Info("Crawl task completed: %s", taskID)
	}()
	logger.Info("start crawl task: %s", taskID)
	response.OkWithMessage(taskID, c)
}
