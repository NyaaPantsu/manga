package controllers

import (
	"github.com/NyaaPantsu/manga/models"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
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
	u := User{}
	if err := c.ParseForm(&u); err != nil {
		//handle error
	}
	_, err := models.GetUserByUsername(u.Username)
	if err != nil {
		password := []byte(u.Password)

		// Hashing the password with the default cost of 10
		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			//
		}
		users := models.Users{
			Username: u.Username,
			Email:    u.Email,
			Password: string(hashedPassword),
		}
		_, err = models.AddUsers(&users)
		if err != nil {

		}
		c.Redirect("/auth/login", 301)
	}

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
