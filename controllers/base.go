package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"
)

// BaseController operations for Base
type BaseController struct {
	beego.Controller

	Userinfo *models.Users
	IsLogin  bool
}
type NestPreparer interface {
	NestPrepare()
}

type NestFinisher interface {
	NestFinish()
}

func (c *BaseController) Prepare() {

	c.IsLogin = c.GetSession("userinfo") != nil
	if c.IsLogin {
		c.Userinfo = c.GetLogin()
	}

	c.Data["IsLogin"] = c.IsLogin
	c.Data["Userinfo"] = c.Userinfo
	c.Layout = "layouts/index.html"
	flash := beego.ReadFromRequest(&c.Controller)
	if n, ok := flash.Data["notice"]; ok {
		// Display settings successful
		c.Data["Notice"] = flash.Data["notice"]
	} else if n, ok = flash.Data["error"]; ok {
		// Display error messages
		c.Data["Error"] = n
	} else if n, ok = flash.Data["warning"]; ok {
		c.Data["Warning"] = flash.Data["warning"]

	} else if n, ok = flash.Data["success"]; ok {
		c.Data["Success"] = flash.Data["success"]
	}
	if app, ok := c.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}

}

func (c *BaseController) Finish() {
	if app, ok := c.AppController.(NestFinisher); ok {
		app.NestFinish()
	}
}

func (c *BaseController) GetLogin() *models.Users {
	u, _ := models.GetUsersById(c.GetSession("userinfo").(int))
	return u
}

func (c *BaseController) DelLogin() {
	c.DelSession("userinfo")
}

func (c *BaseController) SetLogin(user *models.Users) {
	c.SetSession("userinfo", user.Id)
}
