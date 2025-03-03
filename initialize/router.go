package initialize

import (
	"net/http"
	"pacyuribot/global"
	"pacyuribot/middleware"
	"pacyuribot/router"
	"pacyuribot/utils"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()

	// 保证./data/crawl_data文件夹存在
	utils.Mkdir("./data/crawl_data")

	// 保证cache文件夹存在
	utils.Mkdir("./cache")
	r.Static("/cache", "./cache")
	r.Static("/data/crawl_data", "./data/crawl_data")

	AdminGroup := r.Group("/admin")
	UserGroup := r.Group("/user")
	publicGroup := r.Group("/public")

	AdminGroup.Use(middleware.Auth())
	AdminGroup.Use(middleware.DefaultErrorHandler())

	if global.Config.Server.Debug {
		testGroup := r.Group("/test")
		testRouter := router.RouterGroupApp.Test
		testRouter.InitTestCrawlerRouter(testGroup)
		testRouter.InitPocketbaseRouter(testGroup)
	}

	// admin router init
	adminRouter := router.RouterGroupApp.Admin
	adminRouter.InitCrawlerRouter(AdminGroup)
	adminRouter.InitialDatasourceRouter(AdminGroup)

	// user router init
	userRouter := router.RouterGroupApp.User
	userRouter.InitializeChatRouter(UserGroup)
	// 健康监测
	publicGroup.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})
	return r
}
