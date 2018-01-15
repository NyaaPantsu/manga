package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
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

	log := logs.GetLogger()
	log.Println(name)
	flash := beego.NewFlash()
	l, err := models.GetSeriesByName(name)
	log.Println(l.Id)
	if err != nil {
		flash.Error("Comic not found")
		flash.Store(&c.Controller)
		c.Redirect("/comics", 302)
		return
	}
	chap, err := models.GetSeriesChaptersById(l.Id)
	tag, err := models.GetAllTagsById(l.Id)
	c.TplName = "comic.html"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["series"] = l
	c.Data["chapters"] = chap
	c.Data["tags"] = tag
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
