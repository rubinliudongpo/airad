package controllers

// ErrorController definition.
type ErrorController struct {
	BaseController
}

// RetError return error information in JSON.
func (c *ErrorController) RetError(e *ControllerReturn) {
	c.Data["json"] = e
	c.ServeJSON()
}

// Error404 redefine 404 error information.
func (c *ErrorController) Error404() {
	c.RetError(err404)
}