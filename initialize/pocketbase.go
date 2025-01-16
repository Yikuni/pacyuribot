package initialize

import (
	"github.com/pluja/pocketbase"
	"pacyuribot/global"
	"pacyuribot/logger"
)

func InitPocketbase() {
	global.PocketbaseAdminClient = pocketbase.NewClient(
		global.Config.Pocketbase.Url,
		pocketbase.WithAdminEmailPassword(global.Config.Pocketbase.Email, global.Config.Pocketbase.Password),
	)
	_, err := global.Cron.AddFunc("0 0 * * *", RefreshAdminToken)
	if err != nil {
		logger.Error("Failed to add function to Cron: %s", err.Error())
		panic(err.Error())
		return
	}
}

func RefreshAdminToken() {
	pocketbase.WithAdminToken(global.PocketbaseAdminClient.AuthStore().Token())(global.PocketbaseAdminClient)
}
