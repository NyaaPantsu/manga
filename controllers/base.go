package controllers

import (
	"github.com/astaxie/beego"
	"github.com/juusechec/jwt-beego"
	"strings"
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
	auth := c.Ctx.Input.Header("Authorization")
	if len(auth) > 0 {
		header := strings.Split(auth, " ")
		if len(header) != 2 || header[0] != "Bearer" {
			c.Ctx.Output.SetStatus(401)
			c.Data["json"] = "Permission Denied"
			c.ServeJSON()
		}

		et := jwtbeego.EasyToken{}
		valid, _, _ := et.ValidateToken(header[1])
		if !valid {
			c.Ctx.Output.SetStatus(401)
			c.Data["json"] = "Permission Denied"
			c.ServeJSON()
		}
	}
	return

}

func (c *BaseController) Finish() {
	if app, ok := c.AppController.(NestFinisher); ok {
		app.NestFinish()
	}
}
