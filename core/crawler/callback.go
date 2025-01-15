package crawler

import "pacyuribot/logger"

type PageCrawledCallback func(d *DefaultCrawler) bool

func DebugPageCraw(d *DefaultCrawler) bool {
	logger.Info(d.ctx.contentBuilder.Text())
	return false
}
