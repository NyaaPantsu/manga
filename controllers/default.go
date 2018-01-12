package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.tpl"
	c.Layout = "index.tpl"
	c.LayoutSections = make(map[string]string)
	c.Render()
}
