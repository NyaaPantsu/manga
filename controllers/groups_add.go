package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"
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
	Username    string `form:"username,text"`
	Id          int    `form:"-"`
	Name        string `form:"name, text"`
	Description string `form:"description, text"`
	CoverImage  string `form:"cover, file"`
	TypeName    string `form:"typename, text"`
	TypeDemonym string `form:"typede, text"`
	Status      string `form:"status, text"`
	Tags        string `form:"tags, text"`
	Authors     string `form:"author, text"`
	Artist      string `form:"artist, text"`
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

	exists := models.GroupsExists(u.Name)
	if !exists {

		file, header, err := c.GetFile("cover") // where <<this>> is the controller and <<file>> the id of your form field
		if file != nil {
			// get the filename
			fileName := header.Filename
			// save to server
			err := c.SaveToFile("file", "/disk1/covers"+fileName)
			if err != nil {

				flash.Error(err.Error())
				flash.Store(&c.Controller)
				c.Redirect("/groups/add", 301)
				return
			}
		}
		if err != nil {

			flash.Error(err.Error())
			flash.Store(&c.Controller)
			c.Redirect("/groups/add", 301)
			return
		}
		series := models.Groups{}
		_, err = models.AddGroups(&series)
		if err != nil {
			flash.Error(err.Error())
			flash.Store(&c.Controller)
			c.Redirect("/groups/add", 301)
			return
		}
		flash.Success("Successfully added series!")
		flash.Store(&c.Controller)
		c.Redirect("/groups/add", 301)
		return

	}

	flash.Error("Adding series failed")
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
	c.TplName = "series_add.html"
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
