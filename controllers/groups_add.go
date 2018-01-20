package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"
	"github.com/microcosm-cc/bluemonday"

	"github.com/astaxie/beego/utils/pagination"
	"gopkg.in/russross/blackfriday.v2"

	"html/template"
	"strings"
)

// Groups_addController operations for Groups_add
type Groups_addController struct {
	BaseController
}

// URLMapping ...
func (c *Groups_addController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

type GroupsForm struct {
	Id          int    `form:"-"`
	Name        string `form:"name, text"`
	Description string `form:"description, text"`
	ReleaseDlay int    `form:"release_delay, int"`
	Urls        string `form:"urls, text"`
}

// Post ...
// @Title Create
// @Description create Groups
// @Param	body		body 	models.GroupsScanlation	true		"body for Groups_add content"
// @Success 201 {object} models.GroupsScanlation
// @Failure 403 body is empty
// @router / [post]
func (c *Groups_addController) Post() {

	flash := beego.NewFlash()

	u := GroupsForm{}
	if err := c.ParseForm(&u); err != nil {
		flash.Error("Groups invalid")
		flash.Store(&c.Controller)
		c.Redirect("/groups/add", 302)
		return
	}

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

		err := models.AddGroupsScanlation(&groups)
		if err != nil {
			flash.Error(err.Error())
			flash.Store(&c.Controller)
			c.Redirect("/groups/add", 301)
			return
		}
		usergroups := models.UsersGroups{
			GroupName: u.Name,
		}
		err = models.AddUserGroups(&usergroups)
		if err != nil {
			flash.Error(err.Error())
			flash.Store(&c.Controller)
			c.Redirect("/groups/add", 301)
			return
		}

		flash.Success("Successfully added group!")
		flash.Store(&c.Controller)
		c.Redirect("/groups/add", 301)
		return

	}

	flash.Error("Adding a group failed")
	flash.Store(&c.Controller)
	c.Redirect("/groups/add", 302)
	return
}

// Get ...
// @Title Get
// @Description get Groups
// @Success 200 {object} models.GroupsScanlation
// @Failure 403
// @router /add [get]
func (c *Groups_addController) Get() {
	c.TplName = "groups_add.html"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Render()

}

// GetOne ...
// @Title GetOne
// @Description get Groups by name
// @Param	name		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.GroupsScanlation
// @Failure 403 :name is empty
// @router /:name [get]
func (c *Groups_addController) GetOne() {
	name := c.Ctx.Input.Param(":name")

	flash := beego.NewFlash()
	l, err := models.GetGroupsScanlationByName(name)
	if err != nil {
		flash.Error("Comic not found")
		flash.Store(&c.Controller)
		c.Redirect("/groups", 302)
		return
	}
	c.TplName = "group.html"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["group"] = l
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
// @Success 200 {object} models.GroupsScanlation
// @Failure 403
// @router / [get]
func (c *Groups_addController) GetAll() {
	flash := beego.NewFlash()
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 20
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
	if len(order) == 0 {
		order = append(order, "desc")
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

	if len(sortby) == 0 {
		sortby = append(sortby, "name")
	}
	l, err := models.GetAllGroupsScanlation(query, fields, sortby, order, offset, limit)

	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/comics", 302)
		return

	}

	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())

	paginator := pagination.SetPaginator(c.Ctx, int(limit), int64(len(l)))

	c.TplName = "groups.html"
	c.Data["groups"] = l
	c.Data["paginator"] = paginator
	c.Render()
	return

}

// Put ...
// @Title Put
// @Description update the Groups
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.GroupsScanlation	true		"body for Groups content"
// @Success 200 {object} models.GroupsScanlation
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Groups_addController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Groups
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Groups_addController) Delete() {

}
