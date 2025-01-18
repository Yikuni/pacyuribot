package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"pacyuribot/global"
	"pacyuribot/logger"
	"pacyuribot/middleware"
	"pacyuribot/router"
)

func Routers() *gin.Engine {
	r := gin.Default()

	// 保证./data/crawl_data文件夹存在
	_, err := os.Stat("./data/crawl_data")
	if os.IsNotExist(err) {
		// 如果不存在，创建文件夹
		err = os.MkdirAll("./data/crawl_data", os.ModePerm)
		if err != nil {
			logger.Error("Failed to mkdir ./data/crawl_data: %s", err.Error())
			panic(err)
		}
	}
	// 保证cache文件夹存在
	_, err = os.Stat("cache")
	if os.IsNotExist(err) {
		// 如果不存在，创建文件夹
		err = os.Mkdir("cache", os.ModePerm)
		logger.Error("Failed to mkdir ./cache: %s", err.Error())
		if err != nil {
			panic(err)
		}
	}
	r.Static("/cache", "./cache")
	r.Static("/data/crawl_data", "./data/crawl_data")

	AdminGroup := r.Group("/admin")
	//UserGroup := r.Group("/user")
	publicGroup := r.Group("/public")

	AdminGroup.Use(middleware.Auth())

	if global.Config.Server.Debug {
		testGroup := r.Group("/test")
		testRouter := router.RouterGroupApp.Test
		testRouter.InitTestCrawlerRouter(testGroup)
		testRouter.InitPocketbaseRouter(testGroup)
	}

	// admin router init
	adminRouter := router.RouterGroupApp.Admin
	adminRouter.InitCrawlerRouter(AdminGroup)
	// 健康监测
	publicGroup.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})
	return r
}
