package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"

	"github.com/astaxie/beego/validation"

	"html/template"
)

// LoginController operations for Login
type LoginController struct {
	BaseController
}

type Login struct {
	Username string `form:"username,text" valid:"Required"`
	Email    string `form:"email,email"`
	Remember bool   `form:remember, checkbox"`
	Password string `form:"password,password" valid:"Required"`
}

// URLMapping ...
func (c *LoginController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
}

// Post ...
// @Title Create
// @Description create Login
// @Param	body		body 	models.Users	true		"body for Login content"
// @Success 201 {object} models.Users
// @Failure 403 body is empty
// @router / [post]
func (c *LoginController) Post() {
	flash := beego.NewFlash()
	if c.IsLogin {
		flash.Warning("Already logged in")
		flash.Store(&c.Controller)
		c.Redirect("/", 302)
		return
	}
	u := Login{}

	if err := c.ParseForm(&u); err != nil {
		flash.Warning(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/auth/login", 302)
		return
	}

	valid := validation.Validation{}
	b, err := valid.Valid(&u)
	if err != nil {
		flash.Error("Signup invalid!")
		flash.Store(&c.Controller)
		c.Redirect("/auth/signup", 302)
		return
	}
	if !b {
		flash.Error("Signup invalid!")
		flash.Store(&c.Controller)
		c.Redirect("/auth/signup", 302)
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

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), password)
	if err != nil {
		flash.Warning(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/auth/login", 302)
		return

	}
	flash.Success("Successfully logged in")
	flash.Store(&c.Controller)

	c.SetLogin(user)
	// Set the UserID if everything is ok
	c.Redirect("/", 301)

}

// Get ...
// @Title GetAll
// @Description get Login
// @Success 200 {object} models.Users
// @Failure 403
// @router / [get]
func (c *LoginController) Get() {
	c.TplName = "login.html"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Render()
}
