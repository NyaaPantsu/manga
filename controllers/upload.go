package controllers

import (
	"github.com/astaxie/beego"
)

// UploadController operations for Upload
type UploadController struct {
	beego.Controller
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

}

// Get ...
// @Title Get
// @Description get Upload by id
// @Success 200 {object} models.Upload
// @router / [get]
func (c *UploadController) Get() {
	c.Layout = "index.tpl"
	c.TplName = "search.tpl"
	c.Render()
}
