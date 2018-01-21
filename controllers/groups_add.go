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
