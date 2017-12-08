package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

// @router /get_username [GET]
func (this *UserController) GetUsername() {
	result := make(map[string]interface{})
	result["code"] = 1024
	result["msg"] = "success"
	this.Ctx.Output.JSON(result, true, false)
}