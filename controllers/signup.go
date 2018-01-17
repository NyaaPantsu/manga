package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"golang.org/x/crypto/bcrypt"

	"html/template"
)

// SignupController operations for Signup
type SignupController struct {
	BaseController
}

// URLMapping ...
func (c *SignupController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
}

type Signup struct {
	Username  string `form:"username,text" valid:"Required"`
	Email     string `form:"email,email" valid:"Email;Required";`
	Password  string `form:"password,password" valid:"Reqiured"`
	Password2 string `form:"password2, password" valid:"Required"`
}

// Post ...
// @Title Create
// @Description create Signup
// @Param	body		body 	models.Users	true		"body for Signup content"
// @Success 201 {object} models.Users
// @Failure 403 body is empty
// @router / [post]
func (c *SignupController) Post() {
	flash := beego.NewFlash()

	valid := validation.Validation{}
	u := Signup{}
	if err := c.ParseForm(&u); err != nil {
		flash.Error("Signup invalid!")
		flash.Store(&c.Controller)
		c.Redirect("/auth/signup", 302)
		return
	}
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
	username := models.UsernameExists(u.Username)
	email := models.EmailExists(u.Email)
	if u.Password == u.Password2 && !email && !username {
		password := []byte(u.Password)

		// Hashing the password with the default cost of 10
		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			flash.Error(err.Error())
			flash.Store(&c.Controller)
			c.Redirect("/auth/signup", 302)
			return
		}
		users := models.Users{
			Username:     u.Username,
			Email:        u.Email,
			PasswordHash: string(hashedPassword),
		}
		_, err = models.AddUsers(&users)
		if err != nil {
			flash.Error(err.Error())
			flash.Store(&c.Controller)
			c.Redirect("/auth/signup", 302)
			return
		}

		flash.Success("successfully signed up")
		flash.Store(&c.Controller)
		c.Redirect("/auth/login", 301)
		return
	}

	flash.Error("Email or username exists")
	flash.Store(&c.Controller)
	c.Redirect("/auth/login", 302)
	return

}

// Get ...
// @Title Create
// @Description create Signup
// @Param	body		body 	models.Users	true		"body for Signup content"
// @Success 201 {object} models.Users
// @Failure 403 body is empty
// @router / [get]
func (c *SignupController) Get() {
	c.TplName = "signup.html"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Render()

}
