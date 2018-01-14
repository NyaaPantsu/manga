package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"

	"html/template"
)

// Groups_addController operations for Groups_add
type Groups_addController struct {
	BaseController
}

// URLMapping ...
func (c *Groups_addController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

type GroupsForm struct {
	Id          int      `form:"-"`
	Name        string   `form:"name, text"`
	Description string   `form:"description, text"`
	ReleaseDlay int      `form:"release_delay, int"`
	Urls        []string `form:"urls, text"`
}

// Post ...
// @Title Create
// @Description create Groups_add
// @Param	body		body 	models.Groups_add	true		"body for Groups_add content"
// @Success 201 {object} models.Groups_add
// @Failure 403 body is empty
// @router / [post]
func (c *Groups_addController) Post() {

	flash := beego.NewFlash()
	if !c.IsLogin {
		flash.Error("Error you need to be logged in")
		flash.Store(&c.Controller)
		c.Redirect("/", 302)
		return
	}
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
		err := models.AddGroupsScanlation(&groups)
		if err != nil {
			flash.Error(err.Error())
			flash.Store(&c.Controller)
			c.Redirect("/groups/add", 301)
			return
		}
		usergroups := models.UsersGroups{
			GroupName: u.Name,
			UserId:    c.GetLogin(),
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
// @Description get Groups_add
// @Success 200 {object} models.Groups_add
// @Failure 403
// @router / [get]
func (c *Groups_addController) Get() {
	flash := beego.NewFlash()
	if !c.IsLogin {
		flash.Error("Error you need to be logged in")
		flash.Store(&c.Controller)
		c.Redirect("/", 302)
		return
	}
	c.TplName = "groups_add.html"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Render()

}

// Put ...
// @Title Put
// @Description update the Groups_add
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Groups_add	true		"body for Groups_add content"
// @Success 200 {object} models.Groups_add
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Groups_addController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Groups_add
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Groups_addController) Delete() {

}
