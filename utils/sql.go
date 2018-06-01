package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
)

func InitSql() {
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbname := beego.AppConfig.String("db.name")
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8mb4"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dsn)
	if err := orm.RunSyncdb("default", false, true); err != nil {
		fmt.Println(err)
	}
	orm.RunCommand()
}
