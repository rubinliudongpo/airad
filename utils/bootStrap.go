package utils

import (
	"os"
	"syscall"
	"fmt"
	"github.com/astaxie/beego"
	"os/signal"

)

func InitBootStrap()  {
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
	//beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
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

//var FilterUser = func(ctx *context.Context) {
//	_, ok := ctx.Input.Session("userLogin").(string)
//	if !ok && ctx.Request.RequestURI != "/" {
//		ctx.Redirect(302, "/")
//	}
//}

