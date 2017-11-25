package controllers

import (
	"github.com/astaxie/beego"
)

// MainController definition.
type MainController struct {
	beego.Controller
}

// Get method.
func (c *MainController) Get() {
	c.Data["Website"] = "www.liudp.cn"
	c.Data["Email"] = "rubinliu@hotmail.com"
	c.TplName = "index.tpl"
	c.Render()
}
