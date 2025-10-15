package main

import (
	"<%= moduleName %>/bootstrap"
	"<%= moduleName %>/global"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()

	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()

	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()

	// 初始化SQL SERVER数据库
	// global.App.MSDB = bootstrap.InitializeMSDB()

	// 初始化ACCESS数据库
	global.App.MSACCESS = bootstrap.InitializeMSACCESS()

	// 初始化Tgl数据库
	// global.App.TglDB = bootstrap.InitializeTglDB()

	global.App.DispatcherMsgReceiver = make(chan []byte)

	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}

		if global.App.MSDB != nil {
			db, _ := global.App.MSDB.DB()
			db.Close()
		}
	}()

	// 初始化验证器
	bootstrap.InitializeValidator()

	// 初始化Redis
	global.App.Redis = bootstrap.InitializeRedis()

	// 初始化文件系统
	bootstrap.InitializeStorage()

	// 初始化计划任务
	bootstrap.InitializeCron()
	// go ws.GlobalHub.Run()

	// 启动服务器
	bootstrap.RunServer()
}
