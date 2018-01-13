package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"

	"strconv"
)

// ComicController operations for Comic
type ComicController struct {
	beego.Controller
}

// URLMapping ...
func (c *ComicController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
}

// GetOne ...
// @Title GetOne
// @Description get Comic by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Comic
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ComicController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetSeriesById(id)
	if err != nil {

	}
	c.TplName = "series.tpl"
	c.Layout = "index.tpl"
	c.Data["series"] = v
	c.Render()

}

// GetAll ...
// @Title GetAll
// @Description get Comic
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Comic
// @Failure 403
// @router / [get]
func (c *ComicController) GetAll() {

}
