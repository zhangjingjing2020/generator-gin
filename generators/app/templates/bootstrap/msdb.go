package bootstrap

import (
	"<%= moduleName %>/global"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitializeMSDB() *gorm.DB {
	// 根据驱动配置进行初始化
	switch global.App.Config.Msserver.Driver {
	case "sqlserver":
		return initSqlServerGorm()
	default:
		return initSqlServerGorm()
	}
}

func initSqlServerGorm() *gorm.DB {
	dbConfig := global.App.Config.Msserver

	if dbConfig.Database == "" {
		return nil
	}

	dsn := "sqlserver://" + dbConfig.UserName + ":" + dbConfig.Password + "@" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + "?database=" + dbConfig.Database
	if db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,                     // 禁用自动创建外键约束
		Logger:                                   getSqlServerGormLogger(), // 使用自定义 Logger
	}); err != nil {
		global.App.Log.Error("sqlserver connect failed, err:", zap.Any("err", err))
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
		return db
	}
}

func getSqlServerGormLogger() logger.Interface {
	var logMode logger.LogLevel

	switch global.App.Config.Msserver.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}

	return logger.New(getSqlServerGormLogWriter(), logger.Config{
		SlowThreshold:             200 * time.Millisecond,                          // 慢 SQL 阈值
		LogLevel:                  logMode,                                         // 日志级别
		IgnoreRecordNotFoundError: false,                                           // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  !global.App.Config.Msserver.EnableFileLogWriter, // 禁用彩色打印
	})
}

// 自定义 gorm Writer
func getSqlServerGormLogWriter() logger.Writer {
	var writer io.Writer

	// 是否启用日志文件
	if global.App.Config.Msserver.EnableFileLogWriter {
		// 自定义 Writer
		writer = &lumberjack.Logger{
			Filename:   global.App.Config.Log.RootDir + "/" + global.App.Config.Msserver.LogFilename,
			MaxSize:    global.App.Config.Log.MaxSize,
			MaxBackups: global.App.Config.Log.MaxBackups,
			MaxAge:     global.App.Config.Log.MaxAge,
			Compress:   global.App.Config.Log.Compress,
		}
	} else {
		// 默认 Writer
		writer = os.Stdout
	}
	return log.New(writer, "\r\n", log.LstdFlags)
}
