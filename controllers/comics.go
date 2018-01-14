package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"
	"html/template"
)

// ComicsController operations for Comics
type ComicsController struct {
	BaseController
}

// URLMapping ...
func (c *ComicsController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
}

// GetOne ...
// @Title GetOne
// @Description get Comics by name
// @Param	name		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Comics
// @Failure 403 :name is empty
// @router /:name [get]
func (c *ComicsController) GetOne() {
	name := c.Ctx.Input.Param(":name")

	flash := beego.NewFlash()
	l, err := models.GetSeriesByName(name)
	if err != nil {
		flash.Error("Comic not found")
		flash.Store(&c.Controller)
		c.Redirect("/comics", 302)
		return
	}

	c.TplName = "comic.html"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["series"] = l
	c.Render()

}

// GetAll ...
// @Title GetAll
// @Success 200 {object} models.Comics
// @Failure 403
// @router / [get]
func (c *ComicsController) GetAll() {
	c.TplName = "comic.html"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Render()

}
