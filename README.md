## generator-gin

generator-gin 是简单的gin项目工程生成器，常用于Api项目开发

* 支持秒级计划任务
* 支持Mysql、access数据库
* redis缓存
* 日志轮转


### 安装

    npm install -g yo generator-gin

### 生成工程

    yo gin
    
生成的目录结构如下：
```plaintext
.
├── Dockerfile
├── app
│   ├── common
│   │   ├── request
│   │   │   ├── user.go
│   │   │   └── validator.go
│   │   └── response
│   │       └── response.go
│   ├── controllers
│   │   ├── app
│   │   │   └── auth.go
│   │   └── common
│   │       └── upload.go
│   ├── middleware
│   │   ├── cors.go
│   │   ├── jwt.go
│   │   └── recovery.go
│   ├── models
│   │   ├── common.go
│   │   ├── paramConfig.go
│   │   └── user.go
│   ├── services
│   │   ├── jwt.go
│   │   ├── paramConfig.go
│   │   ├── types
│   │   │   ├── pSlice.go
│   │   │   └── pagination.go
│   │   └── user.go
│   └── ws
│       ├── client.go
│       ├── home.go
│       ├── home.html
│       ├── hub.go
│       └── init.go
├── bootstrap
│   ├── accessdb.go
│   ├── config.go
│   ├── cron.go
│   ├── db.go
│   ├── log.go
│   ├── msdb.go
│   ├── redis.go
│   ├── router.go
│   ├── storage.go
│   └── validator.go
├── build_win.bat
├── config
│   ├── app.go
│   ├── config.go
│   ├── database.go
│   ├── jwt.go
│   ├── log.go
│   ├── msaccess.go
│   ├── msserver.go
│   ├── redis.go
│   └── storage.go
├── etc
│   └── config.yaml
├── global
│   ├── app.go
│   ├── error.go
│   └── lock.go
├── go.mod
├── go.mod.tpl
├── go.sum
├── main.go
├── readme.md
├── routes
│   └── api.go
├── static
│   └── dist
│       ├── assets
│       │   ├── index.31b3d229.js
│       │   ├── index.459f8680.css
│       │   ├── logo.03d6d6da.png
│       │   └── vendor.2acfe60d.js
│       ├── favicon.ico
│       └── index.html
├── storage
│   ├── app
│   │   └── public
│   └── logs
│       ├── mssql.log
│       └── sql.log
└── utils
    ├── bcrypt.go
    ├── directory.go
    ├── fmt.go
    ├── http.go
    ├── md5.go
    ├── str.go
    └── validator.go
```

### 运行服务
go mod tidy
go run main.go

* 配置文件在 etc/config.yaml

### 测试用例
