package admin

import (
	"fmt"
	"pacyuribot/global"
	"pacyuribot/response"
	"strings"

	"github.com/pluja/pocketbase"
)

type DatasourceService struct {
}

func (s *DatasourceService) GetDatasource(id string, userID string) (response.DataSource, error) {
	datasource, err := pocketbase.CollectionSet[response.Datasource](global.PocketbaseAdminClient, "data_source").One(datasourceID)
	if err != nil {
		return nil, err
	}
	if !strings.EqualFold(datasource.Owner, userID) {
		return nil, fmt.Errorf("Failed to find datasource: %s", id)
	}
	return datasource, nil
}
