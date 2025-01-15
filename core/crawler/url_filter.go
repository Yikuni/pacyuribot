package crawler

import (
	"net/url"
	"pacyuribot/logger"
	"slices"
	"strings"
)

type UrlFilter func(targetURL *url.URL, d *DefaultCrawler) (*url.URL, bool)

func GetDomainFilter(allowExternalLink bool) UrlFilter {
	return func(targetURL *url.URL, d *DefaultCrawler) (*url.URL, bool) {
		if !slices.Contains(d.DisallowDomains, targetURL.Host) && // 不是不允许的域
			(slices.Contains(d.AllowDomains, targetURL.Host) || // 允许域里面有或者是一次外链
				allowExternalLink && slices.Contains(d.AllowDomains, d.ctx.currentURL.Host)) {
			// 过滤重复的
			for e := d.urlList.Front(); e != nil; e = e.Next() {
				item := e.Value.(UrlListItem)
				if strings.EqualFold(item.targetURL.String(), targetURL.String()) {
					logger.Debug("URL Filter: %s; Reason: 重复URL", targetURL.String())
					return targetURL, true
				}
			}
			return targetURL, false
		}
		logger.Debug("URL Filter: %s; Reason: 不允许或多次外链", targetURL.String())
		return targetURL, true
	}
}

func GetMaxDepthFilter(depth int) UrlFilter {
	return func(targetURL *url.URL, d *DefaultCrawler) (*url.URL, bool) {
		return targetURL, d.ctx.currentDepth >= depth // 0：只允许爬取选中的界面
	}
}
