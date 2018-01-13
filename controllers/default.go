package controllers

import (
	"github.com/astaxie/beego"
)

// MainController operations for main page
type MainController struct {
	beego.Controller
}

// Get ...
// @Title Get
// @Description get Main Page
// @Success 200 {object} models.ChapterGroup
// @Failure 403
// @router / [get
func (c *MainController) Get() {
	c.TplName = "home.tpl"
	c.Layout = "index.tpl"
	c.Render()
}
