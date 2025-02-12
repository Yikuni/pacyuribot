package request

type DefaultCrawlerConfig struct {
	TitleFilter          bool     `json:"title_filter"`
	MaxLengthC           int      `json:"max_length_c"`
	MaxLengthE           int      `json:"max_length_e"`
	MinPageContentLength int      `json:"min_page_content_length"`
	AllowOrigins         []string `json:"allow_origins"`
	DisAllowOrigins      []string `json:"dis_allow_origins"`
	TargetURLS           []string `json:"target_urls"`
	MaxDepth             int      `json:"max_depth"`
	AllowExternalLink    bool     `json:"allow_external_link"`
}

type CrawlerTaskStatus struct {
}
type CrawlTask struct {
	Owner      string               `json:"owner"`
	DataSource string               `json:"data_source"`
	Completed  bool                 `json:"completed"`
	Config     DefaultCrawlerConfig `json:"config"`
}

// CrawlData 文件直接保存在/data/crawl/:id
type CrawlData struct {
	Owner      string `json:"owner"`
	DataSource string `json:"data_source"`
	TargetURL  string `json:"target_url"`
	Modified   bool   `json:"modified"`
	Deleted    bool   `json:"deleted"`
}
