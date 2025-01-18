package crawler

import (
	"os"
	"pacyuribot/logger"
	"pacyuribot/service"
)

type PageCrawledCallback func(d *DefaultCrawler) bool

var (
	s = service.ServiceGroupApp.AdminServiceGroup.CrawlerService
)

func DebugPageCraw(d *DefaultCrawler) bool {
	logger.Info(d.ctx.contentBuilder.Text())
	return false
}

// GetAddCrawlDataCallback 将爬取的数据存到数据库
func GetAddCrawlDataCallback(owner string, datasource string) PageCrawledCallback {
	return func(d *DefaultCrawler) bool {

		dataID, err := s.CreateCrawlData(owner, datasource, d.ctx.currentURL.String())
		if err != nil {
			logger.Error("Failed to Add Crawl Data: %s", err.Error())
			return true
		}
		filePath := "./data/crawl_data/" + dataID + ".txt"
		err = os.WriteFile(filePath, []byte(d.ctx.contentBuilder.Text()), os.ModePerm)
		if err != nil {
			logger.Error("Failed to write file: %s; Reason: %s", filePath, err.Error())
			return true
		}
		return false
	}
}

// GetSmallFileFilter 过滤内容太短的页面
func GetSmallFileFilter(minLen int) PageCrawledCallback {
	return func(d *DefaultCrawler) bool {
		if len(d.ctx.contentBuilder.Text()) < minLen {
			logger.Debug("Filtered page: %s; Reason: content is too short", d.ctx.currentURL)
			return true
		}
		return false
	}
}
