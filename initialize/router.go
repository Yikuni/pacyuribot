package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pacyuribot/global"
	"pacyuribot/router"
)

func Routers() *gin.Engine {
	r := gin.Default()
	r.Static("/cache", "./cache")

	//AdminGroup := r.Group("/admin")
	//UserGroup := r.Group("/user")
	publicGroup := r.Group("/public")

	if global.Config.Server.Debug {
		testGroup := r.Group("/test")
		testRouter := router.RouterGroupApp.Test
		testRouter.InitTestCrawlerRouter(testGroup)
	}

	// 健康监测
	publicGroup.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})
	return r
}
