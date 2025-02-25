package utils

import "fmt"

func GetCrawlFilePath(datasourceID string, id string, fileExtension string) string {
	return fmt.Sprintf("./data/crawl_data/%s/%s.%s", datasourceID, id, fileExtension)
}
