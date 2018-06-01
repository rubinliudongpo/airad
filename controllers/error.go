package controllers

import "github.com/astaxie/beego"

// ErrorController definition.
type ErrorController struct {
	beego.Controller
}


func (c *ErrorController) Error404() {
	c.Data["json"] = Response{
		ErrorCode: 404,
		ErrorMessage:  "Not Found",
	}
	c.ServeJSON()
}
func (c *ErrorController) Error401() {
	c.Data["json"] = Response{
		ErrorCode: 401,
		ErrorMessage:  "Permission denied",
	}
	c.ServeJSON()
}
func (c *ErrorController) Error403() {
	c.Data["json"] = Response{
		ErrorCode: 403,
		ErrorMessage:  "Forbidden",
	}
	c.ServeJSON()
}