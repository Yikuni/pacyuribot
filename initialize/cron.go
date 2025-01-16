package initialize

import (
	"github.com/robfig/cron/v3"
	"pacyuribot/global"
)

func InitializeCron() {
	global.Cron = cron.New()
}
