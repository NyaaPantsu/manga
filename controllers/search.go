package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"

	"strings"
)

// SearchController operations for Search
type SearchController struct {
	BaseController
}

// URLMapping ...
func (c *SearchController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
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
// @Success 200 {object} models.Search
// @Failure 403
// @router / [get]
func (c *SearchController) GetAll() {
	flash := beego.NewFlash()
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

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
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			var k, v string
			if len(kv) != 2 {
				k = "name__icontains"
				v = kv[0]
			} else {

				k, v = kv[0], kv[1]
			}
			query[k] = v
		}
	}

	l, err := models.GetAllSeries(query, fields, sortby, order, offset, limit)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		return

	}
	c.TplName = "search.html"
	c.Data["series"] = l
	c.Render()
}
