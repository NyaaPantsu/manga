package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"
)

// ReaderController operations for Reader
type ReaderController struct {
	beego.Controller
}

// URLMapping ...
func (c *ReaderController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
}

// GetOne ...
// @Title Get One
// @Description get SeriesChapters
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SeriesChapters
// @Failure 403
// @router /:hash [get]
func (c *ReaderController) GetOne() {

	var limit int64 = 50
	var offset int64

	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("p"); err == nil {
		offset = v
	}

	hash := c.Ctx.Input.Param(":hash")
	l, err := models.GetSeriesChaptersByHash(hash)

	if err != nil {

		return
	}
	v, err := models.GetAllChapterFilesById(l.Id, limit, offset)
	if err != nil {

		return
	}

	c.Data["json"] = v

	c.ServeJSON()
	return
}
