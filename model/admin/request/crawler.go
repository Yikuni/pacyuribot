package request

type DefaultCrawlerConfig struct {
	TitleFilter       bool     `json:"title_filter"`
	MaxLengthC        int      `json:"max_length_c"`
	MaxLengthE        int      `json:"max_length_e"`
	AllowOrigins      []string `json:"allow_origins"`
	DisAllowOrigins   []string `json:"dis_allow_origins"`
	TargetURLS        []string `json:"target_urls"`
	MaxDepth          int      `json:"max_depth"`
	AllowExternalLink bool     `json:"allow_external_link"`
}

type CrawlerTaskStatus struct {
}
type CrawlerTask struct {
	Owner            string `json:"owner"`
	DataSource       string `json:"data_source"`
	Completed        string `json:"completed"`
	TotalCrawledURLs int    `json:"total_crawled_urls"`
}

// CrawlData 文件直接保存在/data/crawl/:id
type CrawlData struct {
	Owner      string `json:"owner"`
	DataSource string `json:"data_source"`
	TargetURL  string `json:"target_url"`
	Modified   bool   `json:"modified"`
	Activated  bool   `json:"activated"`
}
