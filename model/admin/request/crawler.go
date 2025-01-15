package request

type DefaultCrawlerConfig struct {
	TitleFilter       bool     `json:"titleFilter"`
	MaxLengthC        int      `json:"maxLengthC"`
	MaxLengthE        int      `json:"maxLengthE"`
	AllowOrigins      []string `json:"allowOrigins"`
	DisAllowOrigins   []string `json:"disAllowOrigins"`
	TargetURLS        []string `json:"targetURLS"`
	MaxDepth          int      `json:"maxDepth"`
	AllowExternalLink bool     `json:"allowExternalLink"`
}

type CrawlerTaskStatus struct {
}
type CrawlerTask struct {
	Config DefaultCrawlerConfig
}
