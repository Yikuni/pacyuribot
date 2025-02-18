package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pluja/pocketbase"
	"pacyuribot/global"
	AdminRes "pacyuribot/model/admin/response"
	"pacyuribot/model/common/response"
	"strings"
)

type DatasourceAPI struct {
}

// ActivateDatasource
// @Summary 上传文件创建datasource后需要创建vector_store
func (a *DatasourceAPI) ActivateDatasource(c *gin.Context) {
	datasourceID, b := c.GetQuery("datasource")
	userID := c.GetString("userID")
	if !b {
		response.InvalidRequestFormat(c)
		return
	}
	datasource, err := pocketbase.CollectionSet[AdminRes.Datasource](global.PocketbaseAdminClient, "data_source").One(datasourceID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if !strings.EqualFold(datasourceID, datasource.ID) || !strings.EqualFold(userID, datasource.Owner) {
		response.FailWithMessage("Failed to find datasource: "+datasourceID, c)
		return
	}
	fileList := []string{}
	if datasource.Crawl {
		list, err := pocketbase.CollectionSet[AdminRes.CrawlData](global.PocketbaseAdminClient, "crawl_data").
			List(pocketbase.ParamsList{
				Page:    1,
				Size:    32767,
				Filters: fmt.Sprintf("owner='%s' && data_source='%s' && deleted=false", userID, datasourceID),
				Sort:    "",
				Expand:  "",
				Fields:  "",
			})
		if err != nil {
			response.Fail(c)
			return
		}
		for i := range list.Items {
			fileList = append(fileList, "data/crawl_data/"+list.Items[i].File)
		}
	} else {
		//list, err := pocketbase.CollectionSet[AdminRes.DataFile](global.PocketbaseAdminClient, "data_file").
		//	List(pocketbase.ParamsList{
		//		Page:    1,
		//		Size:    32767,
		//		Filters: fmt.Sprintf("owner='%s' && data_source='%s' && deleted=false", userID, datasourceID),
		//		Sort:    "",
		//		Expand:  "",
		//		Fields:  "",
		//	})
		//if err != nil {
		//	response.Fail(c)
		//	return
		//}

	}

}

// DeactivateDatasource
// @Summary 删除或禁用datasource后需要删除vector_store
func (a *DatasourceAPI) DeactivateDatasource(c *gin.Context) {}

// UpdateDatasource
// @Summary 在数据源中添加或删除文件后，需要重新更新vector_store
func (a *DatasourceAPI) UpdateDatasource(c *gin.Context) {}
