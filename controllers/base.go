package controllers

import (
	"github.com/astaxie/beego"
	"github.com/juusechec/jwt-beego"
)

// BaseController operations for Base
type BaseController struct {
	beego.Controller
}
type NestPreparer interface {
	NestPrepare()
}

type NestFinisher interface {
	NestFinish()
}

func (c *BaseController) Prepare() {

	//	tokenString := c.Ctx.Input.Query("tokenString")
	tokenString := c.Ctx.Request.Header.Get("Authorization")

	et := jwtbeego.EasyToken{}
	valid, _, _ := et.ValidateToken(tokenString)
	if !valid {
		c.Ctx.Output.SetStatus(401)
		c.Data["json"] = "Permission Denied"
		c.ServeJSON()
	}
	return

}

func (c *BaseController) Finish() {
	if app, ok := c.AppController.(NestFinisher); ok {
		app.NestFinish()
	}
}
