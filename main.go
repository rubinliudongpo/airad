package main


import (
	_ "github.com/rubinliudongpo/airad/routers"
	"github.com/rubinliudongpo/airad/utils"


	"github.com/astaxie/beego"
	"github.com/rubinliudongpo/airad/controllers"
)

func main() {
	utils.InitSql()
	utils.InitTemplate()
	utils.InitCache()
	utils.InitBootStrap()
	beego.ErrorController(&controllers.ErrorController{})

	beego.Run()
}