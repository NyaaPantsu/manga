package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"
	"html/template"
)

// ComicsController operations for Comics
type ComicsController struct {
	BaseController
}

// URLMapping ...
func (c *ComicsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Get", c.Get)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Comics
// @Param	body		body 	models.Comics	true		"body for Comics content"
// @Success 201 {object} models.Comics
// @Failure 403 body is empty
// @router / [post]
func (c *ComicsController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Comics by id
// @Param	id		path 	int	true		"The key for staticblock"
// @Success 200 {object} models.Comics
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ComicsController) GetOne() {

	flash := beego.NewFlash()
	l, err := models.GetSeriesById(1)
	if err != nil {
		flash.Error("Comic not found")
		flash.Store(&c.Controller)
		c.Redirect("/comics", 302)
		return
	}

	c.TplName = "comic.html"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["series"] = l
	c.Render()

}

// Get ...
// @Title Get
// @Success 200 {object} models.Comics
// @Failure 403
// @router / [get]
func (c *ComicsController) Get() {
	c.TplName = "comic.html"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Render()

}

// Put ...
// @Title Put
// @Description update the Comics
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Comics	true		"body for Comics content"
// @Success 200 {object} models.Comics
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ComicsController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Comics
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ComicsController) Delete() {

}
