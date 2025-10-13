package bootstrap

import (
	"<%= moduleName %>/global"

	"github.com/robfig/cron/v3"
)

func InitializeCron() {
	global.App.Cron = cron.New(cron.WithSeconds())

	go func() {
		//
		global.App.Cron.AddFunc("*/10 * * * * *", func() {

		})

		global.App.Cron.Start()
		defer global.App.Cron.Stop()
		select {}
	}()
}
