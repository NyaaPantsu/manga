package controllers

import (
	"github.com/astaxie/beego"
)

// LoginController operations for Login
type LoginController struct {
	beego.Controller
}

// URLMapping ...
func (c *LoginController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
}

// Post ...
// @Title Create
// @Description create Login
// @Param	body		body 	models.Login	true		"body for Login content"
// @Success 201 {object} models.Login
// @Failure 403 body is empty
// @router / [post]
func (c *LoginController) Post() {

}

// Get ...
// @Title GetAll
// @Description get Login
// @Success 200 {object} models.Login
// @Failure 403
// @router / [get]
func (c *LoginController) Get() {
	c.TplName = "login.tpl"
	c.Layout = "index.tpl"
	c.LayoutSections = make(map[string]string)
	c.Render()
}
