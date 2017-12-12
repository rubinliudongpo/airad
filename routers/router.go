// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"airad/controllers"
	"github.com/astaxie/beego"
	//"airad/filters"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.UserController{}, "post:Login")
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user", beego.NSInclude(&controllers.UserController{},),),
		beego.NSNamespace("/airad", beego.NSInclude(&controllers.AirAdController{},),),
		beego.NSNamespace("/device", beego.NSInclude(&controllers.DeviceController{},),),
	)
	beego.AddNamespace(ns)
	//beego.InsertFilter("/permission/list", beego.BeforeRouter, filters.HasPermission)
	beego.Router("/v1/device/getdevicebyuserid", &controllers.DeviceController{}, "POST:GetDevicesByUserId")
}
