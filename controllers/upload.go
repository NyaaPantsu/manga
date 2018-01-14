package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
)

// UploadController operations for Upload
type UploadController struct {
	BaseController
}

// URLMapping ...
func (c *UploadController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
}

// Post ...
// @Title Create
// @Description create Upload
// @Param	body		body 	models.Upload	true		"body for Upload content"
// @Success 201 {object} models.Upload
// @Failure 403 body is empty
// @router / [post]
func (c *UploadController) Post() {
	flash := beego.NewFlash()
	if c.IsLogin {
		flash.Error("Error you must be logged in to upload")
		flash.Store(&c.Controller)
		c.Redirect("/auth/login", 302)
		return
	}
	file, header, err := c.GetFile("file") // where <<this>> is the controller and <<file>> the id of your form field
	if err != nil {
		flash.Warning(err.Error())
		flash.Store(&c.Controller)
		return
	}
	if file != nil {
		// get the filename
		fileName := header.Filename
		// save to server
		err := c.SaveToFile("file", "/disk1/"+fileName)
		if err != nil {
			flash.Warning(err.Error())
			flash.Store(&c.Controller)
			return
		}
	}

}

// Get ...
// @Title Get
// @Description get Upload by id
// @Success 200 {object} models.Upload
// @router / [get]
func (c *UploadController) Get() {
	c.TplName = "search.html"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Render()
}
