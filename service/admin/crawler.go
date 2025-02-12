package admin

import (
	"github.com/pluja/pocketbase"
	"pacyuribot/global"
	"pacyuribot/model/admin/request"
)

type CrawlerService struct {
}

func (s *CrawlerService) CreateCrawlTask(owner string, dataSource string, config request.DefaultCrawlerConfig) (string, error) {
	resp, err := pocketbase.
		CollectionSet[request.CrawlTask](global.PocketbaseAdminClient, "crawl_task").Create(
		request.CrawlTask{
			Owner:      owner,
			DataSource: dataSource,
			Completed:  false,
			Config:     config,
		},
	)
	if err != nil {
		return "", err
	}
	return resp.ID, err
}

func (s *CrawlerService) CreateCrawlData(owner string, dataSource string, targetURL string) (string, error) {
	resp, err := pocketbase.
		CollectionSet[request.CrawlData](global.PocketbaseAdminClient, "crawl_data").Create(
		request.CrawlData{
			Owner:      owner,
			DataSource: dataSource,
			TargetURL:  targetURL,
			Deleted:    false,
		},
	)
	if err != nil {
		return "", err
	}
	return resp.ID, err
}

func (s *CrawlerService) SetCrawlTaskStatus(id string) error {
	task, err := pocketbase.
		CollectionSet[request.CrawlTask](global.PocketbaseAdminClient, "crawl_task").One(id)
	if err != nil {
		return err
	}
	task.Completed = true
	err = pocketbase.
		CollectionSet[request.CrawlTask](global.PocketbaseAdminClient, "crawl_task").Update(id, task)
	if err != nil {
		return err
	}
	return err
}
