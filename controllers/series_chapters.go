package controllers

import (
	"encoding/json"
	"errors"
	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type Response struct {
	Success  bool          `json:"success"`
	Error    string        `json:"errors,omitempty`
	Response []interface{} `json:"response"`
	Count    int64         `json:"count"`
}

// SeriesChaptersController operations for SeriesChapters
type SeriesChaptersController struct {
	beego.Controller
}

// URLMapping ...
func (c *SeriesChaptersController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create SeriesChapters
// @Param	body		body 	models.SeriesChapters	true		"body for SeriesChapters content"
// @Success 201 {int} models.SeriesChapters
// @Failure 403 body is empty
// @router / [post]
func (c *SeriesChaptersController) Post() {
	var v models.SeriesChapters
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddSeriesChapters(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			var temp []interface{}
			temp = append(temp, v)
			c.Data["json"] = Response{
				Success:  true,
				Response: temp,
				Count:    1,
			}
		} else {
			c.Data["json"] = Response{
				Success: false,
				Error:   err.Error(),
			}
		}
	} else {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}

	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get SeriesChapters by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.SeriesChapters
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SeriesChaptersController) GetOne() {
	id := c.Ctx.Input.Param(":id")
	v, err := models.GetSeriesChaptersByHash(id)
	if err != nil {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}
	} else {
		var temp []interface{}
		temp = append(temp, v)
		c.Data["json"] = Response{
			Success:  true,
			Response: temp,
			Count:    1,
		}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get SeriesChapters
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.SeriesChapters
// @Failure 403
// @router / [get]
func (c *SeriesChaptersController) GetAll() {
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
			if len(kv) != 2 {
				err := errors.New("Error: invalid query key/value pair")
				c.Data["json"] = Response{
					Success: false,
					Error:   err.Error(),
				}
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, count, err := models.GetAllSeriesChapters(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}

	} else {
		c.Data["json"] = Response{
			Success:  true,
			Response: l,
			Count:    count,
		}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the SeriesChapters
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.SeriesChapters	true		"body for SeriesChapters content"
// @Success 200 {object} models.SeriesChapters
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SeriesChaptersController) Put() {

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.SeriesChapters{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateSeriesChaptersById(&v); err == nil {
			c.Data["json"] = Response{
				Success: true,
			}
		} else {
			c.Data["json"] = Response{
				Success: false,
				Error:   err.Error(),
			}
		}
	} else {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the SeriesChapters
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SeriesChaptersController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteSeriesChapters(id); err == nil {
		c.Data["json"] = Response{
			Success: true,
		}
	} else {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}
	}
	c.ServeJSON()
}
