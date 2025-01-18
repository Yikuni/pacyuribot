package crawler

import "pacyuribot/model/admin/request"

func GetDefaultCrawler(config request.DefaultCrawlerConfig, userID string, datasourceID string) Crawler {
	crawler := NewDefaultCrawler()
	crawler.
		AddContentFilter(GetTitleFilter(config.MaxLengthC, config.MaxLengthE), 8).
		AddContentFilter(TrimFilter, 10).
		AddPageCrawledCallback(GetSmallFileFilter(config.MinPageContentLength), 1).
		AddPageCrawledCallback(GetAddCrawlDataCallback(userID, datasourceID), 8).
		AddUrlFilter(GetDomainFilter(config.AllowExternalLink), 9).
		AddUrlFilter(GetMaxDepthFilter(config.MaxDepth), 10).
		AddTargetUrls(config.TargetURLS).
		AddAllowedDomains(config.AllowOrigins).
		AddDisallowedDomains(config.DisAllowOrigins)

	return crawler
}

func GetTestCrawler(config request.DefaultCrawlerConfig) Crawler {
	crawler := NewDefaultCrawler()
	crawler.
		AddContentFilter(GetTitleFilter(config.MaxLengthC, config.MaxLengthE), 8).
		AddContentFilter(TrimFilter, 10).
		AddPageCrawledCallback(DebugPageCraw, 10).
		AddTargetUrls(config.TargetURLS).
		AddUrlFilter(GetDomainFilter(config.AllowExternalLink), 9).
		AddUrlFilter(GetMaxDepthFilter(config.MaxDepth), 10).
		AddAllowedDomains(config.AllowOrigins).
		AddDisallowedDomains(config.DisAllowOrigins)

	return crawler
}
