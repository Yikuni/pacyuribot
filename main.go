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

	// 初始化
	initialize.InitializeCron()
	initialize.InitPocketbase()

	// 创建 Gin 路由
	r := initialize.Routers()

	// 启动服务器
	err := r.Run(":" + strconv.Itoa(global.Config.Server.Port))
	if err != nil {
		return
	}
}
