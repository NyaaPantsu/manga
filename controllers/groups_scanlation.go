package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/NyaaPantsu/manga/utils/auth"
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"
	"github.com/astaxie/beego"

	"strings"
	"encoding/json"
	"errors"
	"net/url"
)

// GroupsScanlationController operations for GroupsScanlation
type GroupsScanlationController struct {
	beego.Controller
}

// URLMapping ...
func (c *GroupsScanlationController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

type Groups struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	ReleaseDelay int    `json:"releasedelay"`
	Urls         string `json:"urls"`
}

// Post ...
// @Title Create
// @Description create Groups
// @Param	body		body 	GroupsForm	true		"body for GroupsScanlation content"
// @Success 201 {object} models.GroupsScanlation
// @Failure 403 body is empty
// @router / [post]
func (c *GroupsScanlationController) Post() {
	_, err := models.GetUserByUsername(auth.GetUsername(c.Ctx))
	if err != nil {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}
		c.ServeJSON()
		return
	}
	u := Groups{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &u); err == nil {

		exists := models.GroupsScanlationNameExists(u.Name)
		if !exists {
			unsafe := blackfriday.Run([]byte(u.Description))
			html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
			groups := models.GroupsScanlation{
				Name:        u.Name,
				Description: string(html),
			}
			// todo prevent xss
			urls := []models.GroupsScanlationUrls{}
			for _, cond := range strings.Split(u.Urls, ",") {
				temp := models.GroupsScanlationUrls{
					GroupName: &groups,
					Url:       cond,
				}
				urls = append(urls, temp)
			}

			if err := models.AddGroupsScanlation(&groups); err == nil {
				models.AddMultiGroupUrl(urls)

				var k []interface{}
				k = append(k, groups)
				k = append(k, urls)
				c.Data["json"] = Response{
					Success:  true,
					Response: k,
					Count:    1,
				}
				c.ServeJSON()
				return

			}
			if err != nil {
				c.Data["json"] = Response{
					Success: false,
					Error:   err.Error(),
				}
				c.ServeJSON()
				return
			}

			if err != nil {
				c.Data["json"] = Response{
					Success: false,
					Error:   err.Error(),
				}
				c.ServeJSON()
				return
			}
		}
	}
	err = errors.New("error unable to add new group")
	c.Data["json"] = Response{
		Success: false,
		Error:   err.Error(),
	}
	c.ServeJSON()
	return
}

// GetOne ...
// @Title Get One
// @Description get GroupsScanlation by name
// @Param	name		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.GroupsScanlation
// @Failure 403 :name is empty
// @router /:name [get]
func (c *GroupsScanlationController) GetOne() {
	name := c.Ctx.Input.Param(":name")
	escaped, _ := url.QueryUnescape(name)
	v, err := models.GetGroupsScanlationByName(escaped)
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
// @Description get GroupsScanlation
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.GroupsScanlation
// @Failure 403
// @router / [get]
func (c *GroupsScanlationController) GetAll() {
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

	l, count, err := models.GetAllGroupsScanlation(query, fields, sortby, order, offset, limit)
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

// @Title Put
// @Description update the GroupsScanlation
// @Param	name		path 	string	true		"The name you want to update"
// @Param	body		body 	models.GroupsScanlation	true		"body for GroupsScanlation content"
// @Success 200 {object} models.GroupsScanlation
// @Failure 403 :name is not int
// @router /:name [put]
func (c *GroupsScanlationController) Put() {
	name := c.Ctx.Input.Param(":name")
	v := models.GroupsScanlation{Name: name}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateGroupsScanlationByName(&v); err == nil {
			c.Data["json"] = "OK"
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
// @Description delete the GroupsScanlation
// @Param	name		path 	string	true		"The name you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 name is empty
// @router /:name [delete]
func (c *GroupsScanlationController) Delete() {
	name := c.Ctx.Input.Param(":name")
	if err := models.DeleteGroupsScanlation(name); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = Response{
			Success: false,
			Error:   err.Error(),
		}
	}
	c.ServeJSON()
}
