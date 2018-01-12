package controllers

import (
	"github.com/astaxie/beego"
)

// SignupController operations for Signup
type SignupController struct {
	beego.Controller
}

// URLMapping ...
func (c *SignupController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
}

type User struct {
	Id       int    `form:"-"`
	Username string `form:"username,text,username:"`
	Email    string `form:"email,email,email:"`
	Password string `form:"password,password,password:"`
}

// Post ...
// @Title Create
// @Description create Signup
// @Param	body		body 	models.Signup	true		"body for Signup content"
// @Success 201 {object} models.Signup
// @Failure 403 body is empty
// @router / [post]
func (c *SignupController) Post() {

}

// Get ...
// @Title Create
// @Description create Signup
// @Param	body		body 	models.Signup	true		"body for Signup content"
// @Success 201 {object} models.Signup
// @Failure 403 body is empty
// @router / [get]
func (c *SignupController) Get() {
	c.TplName = "signup.tpl"
	c.Layout = "index.tpl"
	c.Data["Form"] = &User{}

	c.Render()

}
