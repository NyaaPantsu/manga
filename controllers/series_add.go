package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
)

// Series_addController operations for Series_add
type Series_addController struct {
	beego.Controller
}

// URLMapping ...
func (c *Series_addController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Series_add
// @Param	body		body 	models.Series_add	true		"body for Series_add content"
// @Success 201 {object} models.Series_add
// @Failure 403 body is empty
// @router / [post]
func (c *Series_addController) Post() {

}

// Get ...
// @Title Get
// @Description get Series_add
// @Success 200 {object} models.Series_add
// @Failure 403
// @router / [get]
func (c *Series_addController) Get() {
	c.TplName = "series_add.html"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Render()

}

// Put ...
// @Title Put
// @Description update the Series_add
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Series_add	true		"body for Series_add content"
// @Success 200 {object} models.Series_add
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Series_addController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Series_add
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Series_addController) Delete() {

}
