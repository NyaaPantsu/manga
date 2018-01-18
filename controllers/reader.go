package controllers

import (
	"github.com/NyaaPantsu/manga/models"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/utils/pagination"
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
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
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

	flash := beego.NewFlash()

	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/", 302)
		return
	}
	v, err := models.GetAllChapterFilesById(l.Id, limit, offset)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/", 302)
		return
	}
	query["chapter_language"] = l.ChapterLanguage.Name
	order = append(order, "desc")
	sortby = append(sortby, "title")
	k, err := models.GetAllSeriesChapters(query, fields, sortby, order, offset, limit, l.Id)

	count, _ := models.GetSeriesChaptersFilesCount(l.Id)

	paginator := pagination.SetPaginator(c.Ctx, int(limit), count-1)

	c.Data["paginator"] = paginator
	c.TplName = "reader.html"
	c.Data["files"] = v
	c.Data["chapters"] = k
	c.Render()
}
