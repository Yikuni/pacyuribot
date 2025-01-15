package main

import (
	"github.com/BurntSushi/toml"
	"pacyuribot/global"
	"pacyuribot/initialize"
	"pacyuribot/logger"
	"strconv"
)

func main() {
	// 加载配置
	if _, err := toml.DecodeFile("config.toml", &global.Config); err != nil {
		logger.Fatal("Error decoding TOML file: %s", err)
	}
	logger.DEBUG = global.Config.Server.Debug
	logger.Info("Loaded config")

	// 创建 Gin 路由
	r := initialize.Routers()

	// 启动服务器
	err := r.Run(":" + strconv.Itoa(global.Config.Server.Port))
	if err != nil {
		return
	}
}

func test() {
	//crawler.GetDefaultCrawler(crawler.DefaultCrawlerConfig{
	//	TitleFilter:     true,
	//	MaxLengthC:      10,
	//	MaxLengthE:      18,
	//	AllowOrigins:    []string{"person.zju.edu.cn", "localhost"},
	//	DisAllowOrigins: []string{},
	//}).Run([]string{"https://person.zju.edu.cn/zhaohui"})
	// 创建一个上下文

}
