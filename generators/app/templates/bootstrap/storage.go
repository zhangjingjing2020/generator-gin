package bootstrap

import (
	"<%= moduleName %>/global"

	"github.com/jassue/go-storage/local"
)

func InitializeStorage() {
	_, _ = local.Init(global.App.Config.Storage.Disks.Local)
}
