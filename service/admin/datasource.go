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
type FileOperation func(entity *response.UploadFileEntity) (error, bool)

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

// TraverseAllFiles traverses all files associated with a datasource and applies the provided operation to each file.
func (s *DatasourceService) TraverseAllFiles(datasource *response.Datasource, op FileOperation) error {
	files := list.New()
	if datasource.Crawl {
		fileList, err := getCrawlFiles(datasource)
		if err != nil {
			return err
		}
		for i := range fileList {
			files.PushBack(fileList[i])
		}
	} else {
		fileList, err := getDataFiles(datasource)
		if err != nil {
			return err
		}
		for i := range fileList {
			files.PushBack(fileList[i])
		}
	}

	// 遍历文件
	for e := files.Front(); e != nil; e = e.Next() {
		var fileEntity response.UploadFileEntity
		if datasource.Crawl {
			item := e.Value.(response.CrawlData)
			fileEntity = response.UploadFileEntity{
				ID:      item.ID,
				FileID:  item.FileID,
				Path:    utils.GetCrawlFilePath(datasource.ID, item.ID, item.FileExtension),
				Deleted: item.Deleted,
			}
		} else {
			item := e.Value.(response.DataFile)
			filePath, err := utils.GetPocketbaseFile("data_file", item.ID, item.File)
			if err != nil {
				return err
			}
			fileEntity = response.UploadFileEntity{
				ID:      item.ID,
				FileID:  item.FileID,
				Path:    filePath,
				Deleted: item.Deleted,
			}
		}
		// 执行操作
		err, b := op(&fileEntity)
		if err != nil {
			return err
		}
		// 如果需要更新，更新file记录
		if b {
			if datasource.Crawl {
				data := e.Value.(response.CrawlData)
				data.FileID = fileEntity.FileID
				data.Deleted = fileEntity.Deleted
				err = s.UpdateCrawlFile(&data)
			} else {
				data := e.Value.(response.DataFile)
				data.FileID = fileEntity.FileID
				data.Deleted = fileEntity.Deleted
				err = s.UpdateDataFile(&data)
			}
		}
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

func (s *DatasourceService) UpdateCrawlFile(file *response.CrawlData) error {
	return pocketbase.CollectionSet[response.CrawlData](global.PocketbaseAdminClient, "crawl_data").
		Update(file.ID, *file)

}

func (s *DatasourceService) UpdateDataFile(file *response.DataFile) error {
	return pocketbase.CollectionSet[response.DataFile](global.PocketbaseAdminClient, "data_file").
		Update(file.ID, *file)
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
func getCrawlFiles(datasource *response.Datasource) ([]response.CrawlData, error) {
	crawlDataList, err := pocketbase.CollectionSet[response.CrawlData](global.PocketbaseAdminClient, "crawl_data").
		List(getFileCriteria(datasource.ID))
	if err != nil {
		return nil, err
	}
	return crawlDataList.Items, nil
}
func getDataFiles(datasource *response.Datasource) ([]response.DataFile, error) {
	dataFileList, err := pocketbase.CollectionSet[response.DataFile](global.PocketbaseAdminClient, "data_file").
		List(getFileCriteria(datasource.ID))
	if err != nil {
		return nil, err
	}
	return dataFileList.Items, nil
}
