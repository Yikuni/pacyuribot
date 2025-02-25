package admin

import (
	"github.com/google/uuid"
	"github.com/pluja/pocketbase"
	"pacyuribot/global"
	"pacyuribot/model/admin/request"
	"pacyuribot/model/admin/response"
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

func (s *CrawlerService) CreateCrawlData(owner string,
	dataSource string, targetURL string, fileExtension string) (string, error) {

	uuid.New()
	resp, err := pocketbase.
		CollectionSet[response.CrawlData](global.PocketbaseAdminClient, "crawl_data").Create(
		response.CrawlData{
			Datasource:    dataSource,
			TargetURL:     targetURL,
			Modified:      false,
			FileExtension: fileExtension,
			FileID:        "",
			Owner:         owner,
			Deleted:       false,
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
