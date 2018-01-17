package controllers

import (
	"github.com/astaxie/beego"
)

// LogoutController operations for Logout
type LogoutController struct {
	BaseController
}

// URLMapping ...
func (c *LogoutController) URLMapping() {
	c.Mapping("Get", c.Get)
}

// Get ...
// @Title Logout
// @Description Logout
// @Param	body		body 	models.Users	true		"body for Logout content"
// @Success 201 {object} models.Users
// @Failure 403 body is empty
// @router / [get]
func (c *LogoutController) Get() {
	c.DelLogin()
	flash := beego.NewFlash()
	flash.Success("Successfully logged out")
	flash.Store(&c.Controller)
	c.Redirect("/", 302)
}
