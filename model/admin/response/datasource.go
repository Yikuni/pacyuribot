package response

type Datasource struct {
	Crawl       bool   `json:"crawl"`
	Model       string `json:"model"`
	Name        string `json:"name"`
	VectorStore string `json:"vector_store"`
	ID          string `json:"id"`
	Owner       string `json:"owner"`
	Deleted     bool   `json:"deleted"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
}

type DataFile struct {
	FileName   string `json:"file_name"`
	Datasource string `json:"data_source"`
	FileID     string `json:"file_id"`
	File       string `json:"file"`
	ID         string `json:"id"`
	Owner      string `json:"owner"`
	Deleted    bool   `json:"deleted"`
	Created    string `json:"created"`
	Updated    string `json:"updated"`
}

type CrawlData struct {
	Datasource string `json:"data_source"`
	TargetURL  string `json:"target_url"`
	Modified   string `json:"modified"`
	File       string `json:"file"`
	ID         string `json:"id"`
	Owner      string `json:"owner"`
	Deleted    bool   `json:"deleted"`
	Created    string `json:"created"`
	Updated    string `json:"updated"`
}
