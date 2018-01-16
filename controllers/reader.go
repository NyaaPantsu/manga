package controllers

import (
	"github.com/NyaaPantsu/manga/models"

	"github.com/astaxie/beego"
)

// ReaderController operations for Reader
type ReaderController struct {
	BaseController
}

// URLMapping ...
func (c *ReaderController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
}

// GetOne ...
// @Title GetOne
// @Description get Reader by id
// @Param	hash		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Reader
// @Failure 403 :hash is empty
// @router /:hash [get]
func (c *ReaderController) GetOne() {

	flash := beego.NewFlash()

	hash := c.Ctx.Input.Param(":hash")
	l, err := models.GetSeriesChaptersByHash(hash)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/", 302)
		return
	}
	v, err := models.GetAllChapterFilesById(l.Id)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/", 302)
		return
	}

	c.TplName = "reader.html"
	c.Data["files"] = v
	c.Render()
}
