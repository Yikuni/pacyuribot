package test

import (
	"github.com/BurntSushi/toml"
	"pacyuribot/global"
	"pacyuribot/initialize"
	"pacyuribot/logger"
	"pacyuribot/model/admin/request"
	"pacyuribot/service"
	"testing"
)

var (
	s = service.ServiceGroupApp.AdminServiceGroup.CrawlerService
)

func TestMain(m *testing.M) {
	if _, err := toml.DecodeFile("../config.toml", &global.Config); err != nil {
		logger.Fatal("Error decoding TOML file: %s", err)
	}
	logger.DEBUG = global.Config.Server.Debug
	logger.Info("Loaded config")

	// 初始化
	initialize.InitializeCron()
	initialize.InitPocketbase()
	m.Run()
}

func TestCreateTask(t *testing.T) {
	task, err := s.CreateCrawlTask("ib78b32oie93fqv", "0u981dc5t0r11tg", request.DefaultCrawlerConfig{
		TitleFilter:       false,
		MaxLengthC:        0,
		MaxLengthE:        0,
		AllowOrigins:      []string{"localhost"},
		DisAllowOrigins:   []string{"localhost"},
		TargetURLS:        []string{"localhost"},
		MaxDepth:          0,
		AllowExternalLink: false,
	})
	if err != nil {
		panic(err)
		return
	}
	logger.Info("task: %s", task)
}

func TestSetCrawlTaskStatus(t *testing.T) {
	err := s.SetCrawlTaskStatus("1wkm718bupuxm9o")
	if err != nil {
		panic(err)
		return
	}
}
