package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"

	"html/template"
)

// LoginController operations for Login
type LoginController struct {
	BaseController
}

type Login struct {
	Username string `form:"username,text"`
	Email    string `form:"email,email"`
	Remember bool   `form:remember, checkbox"`
	Password string `form:"password,password"`
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
	flash := beego.NewFlash()
	if c.IsLogin {
		c.Redirect("/", 302)
	}
	u := Login{}

	if err := c.ParseForm(&u); err != nil {
		flash.Warning(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/auth/login", 302)
		return
	}
	user, err := models.GetUserByUsername(u.Username)
	if err != nil {
		flash.Warning(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/auth/login", 302)
		return

	}
	password := []byte(u.Password)
	// Comparing the password with the hash

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), password)
	if err != nil {
		flash.Warning(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/auth/login", 302)
		return

	}
	flash.Success("Success logged in")
	flash.Store(&c.Controller)
	c.SetLogin(user)
	// Set the UserID if everything is ok
	c.Redirect("/", 301)

}

// Get ...
// @Title GetAll
// @Description get Login
// @Success 200 {object} models.Login
// @Failure 403
// @router / [get]
func (c *LoginController) Get() {
	c.TplName = "login.tpl"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Render()
}
