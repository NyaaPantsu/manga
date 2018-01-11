package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	type User struct {
		Id    int         `form:"-"`
		Name  interface{} `form:"username"`
		Age   int         `form:"age,text,age:"`
		Sex   string
		Intro string `form:",textarea"`
	}
	c.TplName = "index.tpl"
	c.Layout = "index.tpl"
	c.LayoutSections = make(map[string]string)
	c.Render()
}
