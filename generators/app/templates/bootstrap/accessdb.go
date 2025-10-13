package bootstrap

import (
	"<%= moduleName %>/global"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/alexbrainman/odbc"
)

func InitializeMSACCESS() *sql.DB {
	// 根据驱动配置进行初始化
	switch global.App.Config.Msserver.Driver {
	case "access":
		return initSqlAccessGorm()
	default:
		return initSqlAccessGorm()
	}
}

func initSqlAccessGorm() *sql.DB {
	dbConfig := global.App.Config.Msaccess

	if dbConfig.Database == "" {
		return nil
	}

	dbfilename := dbConfig.Host + dbConfig.Database
	db, err := sql.Open("odbc", fmt.Sprintf("Driver={Microsoft Access Driver (*.mdb, *.accdb)};Dbq=%s;Uid=%s;Pwd=%s;", dbfilename, dbConfig.UserName, dbConfig.Password))
	if err != nil {
		log.Println("open==>>", err)
	}
	// defer db.Close()

	return db
}
