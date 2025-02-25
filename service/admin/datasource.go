package admin

import (
	"container/list"
	"fmt"
	"pacyuribot/global"
	"pacyuribot/model/admin/response"
	"pacyuribot/utils"
	"strings"

	"github.com/pluja/pocketbase"
)

type DatasourceService struct {
}
type FileOperation func(entity response.UploadFileEntity) error

func (s *DatasourceService) GetDatasource(id string, userID string) (*response.Datasource, error) {
	datasource, err := pocketbase.
		CollectionSet[response.Datasource](global.PocketbaseAdminClient, "data_source").One(id)
	if err != nil {
		return nil, err
	}
	if !strings.EqualFold(datasource.Owner, userID) || datasource.Deleted {
		return nil, fmt.Errorf("failed to find datasource: %s", id)
	}
	return &datasource, nil
}

func (s *DatasourceService) GetAllFiles(datasource *response.Datasource) (*list.List, error) {
	if datasource.Crawl {
		return getCrawlFiles(datasource)
	} else {
		return getDataFiles(datasource)
	}
}

// TraverseAllFiles traverses all files associated with a datasource and applies the provided operation to each file.
func (s *DatasourceService) TraverseAllFiles(datasource *response.Datasource, op FileOperation) error {
	files, err := s.GetAllFiles(datasource)
	if err != nil {
		return err
	}
	for e := files.Front(); e != nil; e = e.Next() {
		fileEntity := e.Value.(response.UploadFileEntity)
		err := op(fileEntity)
		if err != nil {
			return err
		}
	}
	return nil
}
func (s *DatasourceService) UpdateDatasource(datasource *response.Datasource) error {
	return pocketbase.
		CollectionSet[response.Datasource](global.PocketbaseAdminClient, "data_source").
		Update(datasource.ID, *datasource)
}

func (s *DatasourceService) GetAllVectorStore(model *response.Model) ([]string, error) {
	r, err := pocketbase.CollectionSet[response.Datasource](global.PocketbaseAdminClient, "data_source").
		List(pocketbase.ParamsList{
			Page:    1,
			Size:    32767,
			Filters: fmt.Sprintf("deleted=false && model=%s && vector_store!=''", model.ID),
			Sort:    "",
			Expand:  "",
			Fields:  "",
		})
	if err != nil {
		return nil, err
	}
	var vectorStoreList []string
	for i := range r.Items {
		vectorStoreList = append(vectorStoreList, r.Items[i].VectorStore)
	}
	return vectorStoreList, nil
}

func (s *DatasourceService) UpdateModel(model *response.Model) error {
	return pocketbase.
		CollectionSet[response.Model](global.PocketbaseAdminClient, "model").
		Update(model.ID, *model)
}

func (s *DatasourceService) UpdateCrawlFile(file *response.UploadFileEntity) {

}

func (s *DatasourceService) UpdateDataFile(file *response.UploadFileEntity) {

}
func getFileCriteria(datasourceID string) pocketbase.ParamsList {
	FileCriteria := pocketbase.ParamsList{
		Page:    1,
		Size:    32767,
		Filters: fmt.Sprintf("data_source='%s'", datasourceID),
		Sort:    "",
		Expand:  "",
		Fields:  "",
	}
	return FileCriteria
}
func getCrawlFiles(datasource *response.Datasource) (*list.List, error) {
	resultList := list.New()

	crawlDataList, err := pocketbase.CollectionSet[response.CrawlData](global.PocketbaseAdminClient, "crawl_data").
		List(getFileCriteria(datasource.ID))
	if err != nil {
		return nil, err
	}
	for i := range crawlDataList.Items {
		item := crawlDataList.Items[i]
		filePath := utils.GetCrawlFilePath(datasource.ID, item.ID, item.FileExtension)
		resultList.PushBack(response.UploadFileEntity{
			ID:      item.ID,
			FileID:  item.FileID,
			Path:    filePath,
			Deleted: item.Deleted,
		})
	}
	return resultList, nil
}
func getDataFiles(datasource *response.Datasource) (*list.List, error) {
	resultList := list.New()

	dataFileList, err := pocketbase.CollectionSet[response.DataFile](global.PocketbaseAdminClient, "data_file").
		List(getFileCriteria(datasource.ID))
	if err != nil {
		return nil, err
	}
	for i := range dataFileList.Items {
		item := dataFileList.Items[i]
		filePath, err := utils.GetPocketbaseFile("data_file", item.ID, item.File)
		if err != nil {
			return nil, err
		}
		resultList.PushBack(response.UploadFileEntity{
			ID:      item.ID,
			FileID:  item.FileID,
			Path:    filePath,
			Deleted: item.Deleted,
		})
	}
	return resultList, nil
}
