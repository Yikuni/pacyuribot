package global

import (
	"github.com/pluja/pocketbase"
	"github.com/robfig/cron/v3"
	"pacyuribot/config"
)

var (
	Config                config.Config
	PocketbaseAdminClient *pocketbase.Client
	Cron                  *cron.Cron
)
