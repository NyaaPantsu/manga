package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils/pagination"
	"html/template"
	"strconv"
	"strings"
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
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Series
// @Failure 403 :name is empty
// @router /:id/:name [get]
func (c *ComicsController) GetOne() {
	flash := beego.NewFlash()

	id := c.Ctx.Input.Param(":id")

	log := logs.GetLogger()
	i1, err := strconv.Atoi(id)
	log.Println(i1)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/comics", 302)
		return
	}
	l, err := models.GetSeriesById(i1)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/comics", 302)
		return
	}
	chap, err := models.GetSeriesChaptersBySeriesId(l.Id)
	tag, err := models.GetAllTagsById(l.Id)
	c.TplName = "comic.html"
	paginator := pagination.SetPaginator(c.Ctx, 20, int64(len(chap)))

	c.Data["paginator"] = paginator
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["series"] = l
	c.Data["description"] = template.HTML(l.Description)
	c.Data["chapters"] = chap
	c.Data["tags"] = tag
	c.Render()

}

// GetAll ...
// @Title GetAll
// @Description get Search
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Series
// @Failure 403
// @router / [get]
func (c *ComicsController) GetAll() {
	flash := beego.NewFlash()
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 20
	var offset int64
	var series bool
	series, err := c.GetBool("series")
	if err != nil {
		series = false
	}

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	if len(order) == 0 {
		order = append(order, "desc")
	}

	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			var k, v string
			if len(kv) != 2 {
				if series {
					k = "name__icontains"
				} else {
					k = "title__icontains"

				}
				v = kv[0]
			} else {

				k, v = kv[0], kv[1]
			}
			query[k] = v
		}
	}
	var l []interface{}
	if series {

		if len(sortby) == 0 {
			sortby = append(sortby, "name")
		}
		l, err = models.GetAllSeries(query, fields, sortby, order, offset, limit)

		c.TplName = "series.html"
	} else {
		if len(sortby) == 0 {
			sortby = append(sortby, "time_uploaded")
		}
		l, err = models.GetAllSeriesChapters(query, fields, sortby, order, offset, limit, 0)

		c.TplName = "comics.html"
	}
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/comics", 302)
		return

	}

	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())

	paginator := pagination.SetPaginator(c.Ctx, int(limit), int64(len(l)))

	c.Data["series"] = l
	c.Data["paginator"] = paginator
	c.Render()
	return
}
