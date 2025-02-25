package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pluja/pocketbase"
	"pacyuribot/global"
	"pacyuribot/logger"
	adminRes "pacyuribot/model/admin/response"
	"pacyuribot/model/common/response"
)

type DatasourceAPI struct {
}

// ActivateDatasource
// @Summary 上传文件创建datasource后需要创建vector_store，或者文件删除/增加后
// @Security 	Auth
// @Param 		datasource datasourceID
// @Success 	200 {string} string "{"code":0,"data":{},"msg":"成功"}"
// @Router	/admin/datasource/activate
func (a *DatasourceAPI) ActivateDatasource(c *gin.Context) {
	datasourceID, b := c.GetQuery("datasource")
	userID := c.GetString("userID")
	if !b {
		response.InvalidRequestFormat(c)
		return
	}

	// 获取datasource下的所有文件
	datasource, err := datasourceService.GetDatasource(datasourceID, userID)
	if err != nil {
		panic(err)
	}
	// 上传文件或者删除文件，并更改vector store
	var fileList []string
	err = datasourceService.TraverseAllFiles(datasource, func(e *adminRes.UploadFileEntity) (error, bool) {
		if e.FileID != "" {
			if e.Deleted {
				// 删除文件
				err := assistantService.DeleteFile(e.FileID)
				if err != nil {
					e.FileID = ""
				}
				return err, false
			} else {
				// 已经上传过并且没删除
				fileList = append(fileList, e.FileID)
				return nil, false
			}
		} else {
			// 如果文件没上传过
			fileID, err := assistantService.UploadFile(e.Path, getFileName(datasourceID, e.ID))
			// 更新fileID
			if err != nil {
				logger.Error(err.Error())
				return err, false
			}
			fileList = append(fileList, fileID)
			e.FileID = fileID
		}
		return nil, true
	})
	if err != nil {
		panic(err)
	}
	// 删除旧的vector store
	if datasource.VectorStore != "" {
		err = assistantService.DeleteVectorStore(datasource.VectorStore)
		if err != nil {
			panic(err)
		}
	}
	// 创建vector store
	store, err := assistantService.CreateVectorStore(getVectorStoreName(datasourceID), fileList)
	if err != nil {
		panic(err)
	}
	// 将vector store id保存到datasource实体中
	datasource.VectorStore = store
	err = datasourceService.UpdateDatasource(datasource)
	if err != nil {
		panic(err)
	}
	err = updateAssistant(datasource)
	if err != nil {
		panic(err)
	}
	response.Ok(c)
}

// DeactivateDatasource
// @Summary 删除或禁用datasource后需要删除vector_store
// @Router	/admin/datasource/deactivate
func (a *DatasourceAPI) DeactivateDatasource(c *gin.Context) {
	datasourceID, b := c.GetQuery("datasource")
	userID := c.GetString("userID")
	if !b {
		response.InvalidRequestFormat(c)
		return
	}
	// 仅删除vector store，不删除相关文件
	datasource, _ := datasourceService.GetDatasource(datasourceID, userID)
	_ = assistantService.DeleteVectorStore(datasource.VectorStore)
	datasource.VectorStore = ""
	_ = datasourceService.UpdateDatasource(datasource)

	response.Ok(c)
}

func getVectorStoreName(datasourceID string) string {
	return fmt.Sprintf("pacyuri_datasource-%s", datasourceID)
}
func getFileName(datasourceID string, file string) string {
	return fmt.Sprintf("pacyuri_file_%s_%s", datasourceID, file)

}

func updateAssistant(datasource *adminRes.Datasource) error {
	// 获取model
	model, err := pocketbase.CollectionSet[adminRes.Model](global.PocketbaseAdminClient, "models").
		One(datasource.Model)
	if err != nil {
		return err
	}
	// 获取所有vector store并更新assistant
	allVectorStore, err := datasourceService.GetAllVectorStore(&model)
	if err != nil {
		return err

	}

	if model.Assistant != "" {
		err = assistantService.ModifyAssistant(model.Assistant, allVectorStore)
		if err != nil {
			return err
		}
	} else {
		assistant, err := assistantService.CreateAssistant("pacyuri_assistant_"+model.ID, allVectorStore)
		if err != nil {
			return err
		}
		model.Assistant = assistant
		// 将assistant的ID保存到model中
		err = datasourceService.UpdateModel(&model)
		if err != nil {
			return err
		}
	}
	return nil
}
