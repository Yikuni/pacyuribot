package admin

import "github.com/gin-gonic/gin"

type DatasourceAPI struct {
}

// ActivateDatasource
// @Summary 上传文件创建datasource后需要创建vector_store
func (a *DatasourceAPI) ActivateDatasource(c *gin.Context) {}

func (a *DatasourceAPI) ActivateAllDatasource(c *gin.Context) {}

// DeactivateDatasource
// @Summary 删除或禁用datasource后需要删除vector_store
func (a *DatasourceAPI) DeactivateDatasource(c *gin.Context) {}

// UpdateDatasource
// @Summary 在数据源中添加或删除文件后，需要重新更新vector_store
func (a *DatasourceAPI) UpdateDatasource(c *gin.Context) {}
