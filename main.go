package main

import (
	_ "airad-app-api/routers"
	"airad-app-api/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/signal"
	"syscall"
	"fmt"
	//"gopkg.in/mgo.v2"
	//"github.com/astaxie/beego/session"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "gouser:gopassword@tcp(127.0.0.1:3306)/airad?charset=utf8mb4")
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
	}
	orm.RunCommand()
}

func handleSignals(c chan os.Signal) {
	switch <-c {
	case syscall.SIGINT, syscall.SIGTERM:
		fmt.Println("Shutdown quickly, bye...")
	case syscall.SIGQUIT:
		fmt.Println("Shutdown gracefully, bye...")
		// do graceful shutdown
	}

	os.Exit(0)
}

func main() {
	graceful, _ := beego.AppConfig.Bool("Graceful")
	if !graceful {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		go handleSignals(sigs)
	}
	beego.SetLogger("file", `{"filename":"logs/logs.log"}`)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		beego.SetLevel(beego.LevelDebug)
	} else if (beego.BConfig.RunMode == "prod") {
		beego.SetLevel(beego.LevelInformational)
	}
	beego.ErrorController(&controllers.ErrorController{})

	beego.Run()
}
