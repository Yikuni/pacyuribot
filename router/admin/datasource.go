package admin

import "github.com/gin-gonic/gin"

type DatasourceRouter struct {
}

func (d *DatasourceRouter) InitialDatasourceRouter(rGroup *gin.RouterGroup) {
	r := rGroup.Group("datasource")
	r.POST("activate", datasourceAPI.ActivateDatasource)
	r.POST("deactivate", datasourceAPI.DeactivateDatasource)
}
