package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

// Controller Response is controller error info struct.
type Response struct {
	Status int `json:"status"`
	ErrorCode int `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Data interface{} `json:"data"`
}