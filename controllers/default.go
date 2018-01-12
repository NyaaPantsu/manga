package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "home.tpl"
	c.Layout = "index.tpl"
	c.LayoutSections = make(map[string]string)
	c.Render()
}
