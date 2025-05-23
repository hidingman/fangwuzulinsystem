package main

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"go-schema/routers"
	"go-schema/utils"
)

var cfg *ini.File

func init() {
	var err error
	cfg, err = ini.Load("conf/app.conf")
	if err != nil {
		panic(err)
	}
	sqlconn := cfg.Section("").Key("sqlconn").String()
	orm.RegisterDataBase("default", "mysql", sqlconn)
	orm.Debug = false
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := routers.SetupRouter()
	utils.JiaoYangPrint()
	r.Run()
}
